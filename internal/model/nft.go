package model

import (
	"marketplace-backend/pkg/elastic"
	"time"
)

type NFTIndex struct {
	Id              string     `json:"id,omitempty"`
	ContractAddress string     `json:"contract_address,omitempty"`
	Owner           string     `json:"owner,omitempty"`
	Name            string     `json:"name,omitempty"`
	Description     string     `json:"description,omitempty"`
	Price           *NFTPrice  `json:"price,omitempty"`
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

type NFTPrice struct {
	ETH  int  `json:"eth"`
	VNGT uint `json:"vngt"`
}

type NFTSearchReq struct {
	ResponseConfig elastic.ResponseSearchConfig `json:"responseConfig"`
	Text           string                       `json:"text"`
	Attrs          []NFTAttrsSearchReq          `json:"attrs"`
	SaleType       []string                     `json:"saleType"`
	Price          NFTPriceSearchReq            `json:"price"`
}

type NFTPriceSearchReq struct {
	Currency string                `json:"currency"`
	Range    elastic.RangeQueryReq `json:"range"`
}

type NFTAttrsSearchReq struct {
	TraitType    string   `json:"traitType"`
	DisplayValue []string `json:"displayValue"`
	Value        []string `json:"value"`
}
