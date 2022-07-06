package vmp

import (
	"fmt"
	"marketplace-backend/pkg/elasticstore"

	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

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
	ResponseConfig elasticstore.ResponseSearchConfig `json:"responseConfig"`
	Text           string                            `json:"text"`
	Attrs          []AttrsReq                        `json:"attrs"`
	SaleType       []string                          `json:"saleType"`
	Price          elasticstore.RangeQueryReq        `json:"price"`
}

type AttrsReq struct {
	TraitType    string   `json:"traitType"`
	DisplayValue []string `json:"displayValue"`
}

type SearchResp struct {
	Data any `json:"data"`
	Req  any `json:"req"`
}

func (h *Handler) SearchNFT(r *http.Request, req *SearchNFTRequest, resp *elasticstore.SearchResults[NFTIndex]) error {
	store := elasticstore.NewStore[NFTIndex](h.esClient, "marketplace-nfts")
	var must []interface{}
	if req.Text == "" {
		must = append(must, store.BuildMatchAllQuery())
	} else {
		must = append(must, store.BuildMultiMatchQuery(req.Text, []string{"name", "description"}, true, 2))
	}

	if len(req.SaleType) > 0 {
		must = append(must, store.BuildTermsQuery("sale_type", req.SaleType))
	}
	if (elasticstore.RangeQueryReq{}) != req.Price {
		mRange := store.BuildRangeQuery(&req.Price, "price")
		must = append(must, mRange)
	}

	if len(req.Attrs) > 0 {
		for _, attr := range req.Attrs {
			var nestedMust []interface{}
			if attr.TraitType != "" {
				nestedMust = append(nestedMust, store.BuildTermQuery("attributes.trait_type", attr.TraitType))
			}
			if len(attr.DisplayValue) > 0 {

				nestedMust = append(nestedMust, store.BuildTermsQuery("attributes.display_value", attr.DisplayValue))
			}

			must = append(must, store.BuildNestedQuery(
				"attributes",
				store.BuildBoolQuery("must", &nestedMust),
			))
		}

	}

	mapQuery := store.BuildBoolQuery("must", &must)

	if req.ResponseConfig.Size < 1 {
		return fmt.Errorf("limit: min=1")
	}

	if req.ResponseConfig.From < 0 {
		return fmt.Errorf("offset: min=0")
	}
	store.SetResponseSearchConfig(req.ResponseConfig)
	queryBuild, err := store.BuildQuery(mapQuery)
	if err != nil {
		return err
	}

	err = store.SearchByQuery(queryBuild, resp)

	if err != nil {
		log.Err(err).Msg("elasticsearch error")
		return fmt.Errorf("Internal server error")
	}
	return nil
}
