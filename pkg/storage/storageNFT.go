package storage

import (
	"context"
	"marketplace-backend/pkg/elastic"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type StorageNFTSrv interface {
	InsertNFT(ctx context.Context, NFT NFTIndex, docId string) error
	UpdateNFT(ctx context.Context, NFT NFTIndex, docId string) error
	DeleteNFT(ctx context.Context, id string) error
	FindNFTById(ctx context.Context, id string) (*NFTIndex, error)
	SearchByQuery(ctx context.Context, nft SearchNFTRequest) (elastic.SearchResults[NFTIndex], error)
}

type StorageNFT struct {
	es        *elasticsearch.Client
	indexName string
	timeout   time.Duration
	storeSrv  elastic.StoreSrv[NFTIndex]
}

func NewStorageNFTSrv(es *elasticsearch.Client, timeout time.Duration) StorageNFTSrv {
	indexName := "marketplace-nfts"
	storeSrv := elastic.NewStoreSrv[NFTIndex](es, indexName)
	return &StorageNFT{
		es:        es,
		timeout:   timeout,
		indexName: indexName,
		storeSrv:  storeSrv,
	}
}

type NFTIndex struct {
	Id              string     `json:"id,omitempty"`
	ContractAddress string     `json:"contract_address,omitempty"`
	Owner           string     `json:"owner,omitempty"`
	Name            string     `json:"name,omitempty"`
	Description     string     `json:"description,omitempty"`
	Price           int        `json:"price,omitempty"`
	SaleType        string     `json:"sale_type,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	CreatedTime     *time.Time `json:"created_time,omitempty"`
	LastSoldTime    *time.Time `json:"last_sold_time,omitempty"`
	ListedTime      *time.Time `json:"listed_time,omitempty"`
	Attributes      []NFTAttrs `json:"attributes,omitempty"`
}

type NFTAttrs struct {
	TraitType    string `json:"trait_type,omitempty"`
	DisplayValue string `json:"display_value,omitempty"`
	Value        string `json:"value,omitempty"`
}

type SearchNFTRequest struct {
	ResponseConfig elastic.ResponseSearchConfig `json:"responseConfig"`
	Text           string                       `json:"text"`
	Attrs          []AttrsReq                   `json:"attrs"`
	SaleType       []string                     `json:"saleType"`
	Price          elastic.RangeQueryReq        `json:"price"`
}

type AttrsReq struct {
	TraitType    string   `json:"traitType"`
	DisplayValue []string `json:"displayValue"`
	Value        []string `json:"value"`
}

func (store *StorageNFT) InsertNFT(ctx context.Context, NFT NFTIndex, docId string) error {
	storeSrv := store.storeSrv
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	err := storeSrv.CreateIndex(ctx, NFT, docId)
	if err != nil {
		return err
	}
	return nil
}

func (store *StorageNFT) UpdateNFT(ctx context.Context, NFT NFTIndex, docId string) error {
	storeSrv := store.storeSrv
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	err := storeSrv.UpdateIndex(ctx, NFT, docId)
	if err != nil {
		return err
	}
	return nil
}

func (store *StorageNFT) DeleteNFT(ctx context.Context, docId string) error {
	storeSrv := store.storeSrv
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	err := storeSrv.DeleteIndex(ctx, docId)
	if err != nil {
		return err
	}
	return nil
}

func (store *StorageNFT) FindNFTById(ctx context.Context, id string) (*NFTIndex, error) {
	return nil, nil
}

func (store *StorageNFT) SearchByQuery(ctx context.Context, req SearchNFTRequest) (elastic.SearchResults[NFTIndex], error) {
	result := elastic.SearchResults[NFTIndex]{}
	storeSrv := store.storeSrv
	var must []interface{}

	if req.Text == "" {
		must = append(must, storeSrv.BuildMatchAllQuery())
	} else {
		must = append(must, storeSrv.BuildMultiMatchQuery(req.Text, []string{"name", "description"}, true, 2))
	}

	if len(req.SaleType) > 0 {
		must = append(must, storeSrv.BuildTermsQuery("sale_type", req.SaleType))
	}
	if (elastic.RangeQueryReq{}) != req.Price {
		mRange := storeSrv.BuildRangeQuery(&req.Price, "price")
		must = append(must, mRange)
	}

	if len(req.Attrs) > 0 {
		for _, attr := range req.Attrs {
			var nestedMust []interface{}
			if attr.TraitType != "" {
				nestedMust = append(nestedMust, storeSrv.BuildTermQuery("attributes.trait_type", attr.TraitType))
			}
			if len(attr.DisplayValue) > 0 {

				nestedMust = append(nestedMust, storeSrv.BuildTermsQuery("attributes.display_value", attr.DisplayValue))
			}
			if len(attr.Value) > 0 {
				nestedMust = append(nestedMust, storeSrv.BuildTermsQuery("attributes.value", attr.Value))
			}

			must = append(must, storeSrv.BuildNestedQuery(
				"attributes",
				storeSrv.BuildBoolQuery("must", &nestedMust),
			))
		}

	}

	mapQuery := storeSrv.BuildBoolQuery("must", &must)

	storeSrv.SetResponseSearchConfig(req.ResponseConfig)
	queryBuild, err := storeSrv.BuildQuery(mapQuery)
	if err != nil {
		return result, err
	}
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	result, err = storeSrv.SearchByQuery(ctx, queryBuild)
	if err != nil {
		return result, err
	}

	return result, nil
}
