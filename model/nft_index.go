package model

import (
	"time"
)

type NFTIndex struct {
	TokenId         string                 `json:"token_id,omitempty"`
	ContractAddress string                 `json:"contract_address,omitempty"`
	Owner           string                 `json:"owner,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Price           map[string]interface{} `json:"price,omitempty"`
	SaleType        string                 `json:"sale_type,omitempty"`
	UpdatedAt       *time.Time             `json:"updated_at,omitempty"`
	CreatedTime     *time.Time             `json:"created_time,omitempty"`
	LastSoldTime    *time.Time             `json:"last_sold_time,omitempty"`
	ListedTime      *time.Time             `json:"listed_time,omitempty"`
	Attributes      any                    `json:"attributes,omitempty"`
}

type NFTIndexSearchReq struct {
	ResponseConfig ResponseSearchConfig     `json:"responseConfig"`
	Text           string                   `json:"text"`
	Attrs          []NFTIndexAttrsSearchReq `json:"attrs"`
	SaleType       []string                 `json:"saleType"`
	Price          NFTIndexPriceSearchReq   `json:"price"`
}

type NFTIndexPriceSearchReq struct {
	Currency string        `json:"currency"`
	Range    RangeQueryReq `json:"range"`
}

type NFTIndexAttrsSearchReq struct {
	TraitType    string   `json:"traitType"`
	DisplayValue []string `json:"displayValue"`
	Value        []string `json:"value"`
}

type NFTIndexSearchRes struct {
	Pagination *Paging  `json:"pagination"`
	Data       *[]Token `json:"data,omitempty"`
}
