package storage

import (
	"context"
	"fmt"
	"marketplace-backend/model"
	"marketplace-backend/pkg/elastic"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type StorageNFTSrv interface {
	InsertNFT(ctx context.Context, NFT model.NFTIndex, docId string) error
	UpdateNFT(ctx context.Context, NFT model.NFTIndex, docId string) error
	DeleteNFT(ctx context.Context, token model.Token) error
	SearchByQuery(ctx context.Context, nft model.NFTIndexSearchReq) (model.SearchResults, error)
}

type StorageNFT struct {
	es        *elasticsearch.Client
	indexName string
	timeout   time.Duration
	storeSrv  elastic.StoreSrv[model.NFTIndex]
}

func NewStorageNFTSrv(es *elasticsearch.Client, timeout time.Duration) StorageNFTSrv {
	indexName := "marketplace-nfts"
	storeSrv := elastic.NewStoreSrv[model.NFTIndex](es, indexName)
	return &StorageNFT{
		es:        es,
		timeout:   timeout,
		indexName: indexName,
		storeSrv:  storeSrv,
	}
}

func (store *StorageNFT) InsertNFT(ctx context.Context, NFT model.NFTIndex, docId string) error {
	storeSrv := store.storeSrv
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	err := storeSrv.InsertIndex(ctx, NFT, docId)
	if err != nil {
		return err
	}
	return nil
}

func (store *StorageNFT) UpdateNFT(ctx context.Context, NFT model.NFTIndex, docId string) error {
	storeSrv := store.storeSrv
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	err := storeSrv.UpdateIndex(ctx, NFT, docId)
	if err != nil {
		return err
	}
	return nil
}

func (store *StorageNFT) SearchByQuery(ctx context.Context, req model.NFTIndexSearchReq) (model.SearchResults, error) {
	result := model.SearchResults{}
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

	if (model.NFTIndexPriceSearchReq{} != req.Price) && (model.RangeQueryReq{} != req.Price.Range) {
		mRange := storeSrv.BuildRangeQuery(&req.Price.Range, fmt.Sprintf("price.%v", req.Price.Currency))
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
	queryBuild, err := storeSrv.BuildSearchQuery(mapQuery)
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

func (store *StorageNFT) DeleteNFT(ctx context.Context, token model.Token) error {
	storeSrv := store.storeSrv
	var must []interface{}
	must = append(
		must,
		storeSrv.BuildTermQuery("owner", token.Owner),
		storeSrv.BuildTermQuery("contract_address", token.ContractAddress),
		storeSrv.BuildTermQuery("token_id", token.TokenID),
	)
	mapQuery := storeSrv.BuildBoolQuery("must", &must)
	queryBuild, err := storeSrv.BuildModifyQuery(mapQuery)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	err = storeSrv.DeleteIndexByQuery(ctx, queryBuild)
	if err != nil {
		return err
	}

	return nil
}
