package storage

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"marketplace-backend/pkg/elastic"
	"time"
)

type NFTStorageSrv interface {
	Insert(ctx context.Context, NFT NFTIndex) error
	Update(ctx context.Context, NFT NFTIndex) error
	Delete(ctx context.Context, id string) error
	FindOne(ctx context.Context, id string) (*NFTIndex, error)
}

type NFTStorage struct {
	es        *elasticsearch.Client
	indexName string
	timeout   time.Duration
}

func NewNFTSrv(es *elasticsearch.Client, timeout time.Duration) NFTStorageSrv {
	return &NFTStorage{
		es:        es,
		timeout:   timeout,
		indexName: "marketplace-nfts",
	}
}

type NFTIndex struct {
	NftId           string     `json:"nft_id,omitempty"`
	CollectionId    string     `json:"collection_id,omitempty"`
	Name            string     `json:"name,omitempty"`
	Description     string     `json:"description,omitempty"`
	Price           int        `json:"price,omitempty"`
	SaleType        string     `json:"sale_type,omitempty"`
	CreatedTime     *time.Time `json:"created_time,omitempty"`
	LastSoldTime    *time.Time `json:"last_sold_time,omitempty"`
	ListedTime      *time.Time `json:"listed_time,omitempty"`
	BackgroundColor string     `json:"background_color,omitempty"`
	Image           string     `json:"image,omitempty"`
	Attributes      []NFTAttrs `json:"attributes,omitempty"`
}

type NFTAttrs struct {
	TraitType    string `json:"trait_type,omitempty"`
	DisplayValue string `json:"display_value,omitempty"`
	Value        int    `json:"value,omitempty"`
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
}

func (store *NFTStorage) Insert(ctx context.Context, NFT NFTIndex) error {
	return nil
}

func (store *NFTStorage) Update(ctx context.Context, NFT NFTIndex) error {
	return nil
}

func (store *NFTStorage) Delete(ctx context.Context, id string) error {
	return nil
}

func (store *NFTStorage) FindOne(ctx context.Context, id string) (*NFTIndex, error) {
	return nil, nil
}

func (store *NFTStorage) SearchByQuery(ctx context.Context, id string) (*elastic.SearchResults[NFTIndex], error) {
	result := elastic.SearchResults[NFTIndex]{}

	return &result, nil
}
