package model

import (
	"marketplace-backend/pkg/elastic"
	"time"
)

type Token struct {
	ObjId           string                 `json:"_id" bson:"_id"`
	ContractAddress string                 `json:"contract_address,omitempty" bson:"contract_address"`
	ID              string                 `json:"id,omitempty" bson:"id"`
	Owner           string                 `json:"owner,omitempty" bson:"owner"`
	PreviousOwner   string                 `json:"-" bson:"previous_owner"`
	Metadata        map[string]interface{} `json:"metadata,omitempty" bson:"metadata"`
	MetadataRaw     string                 `json:"-" bson:"metadata_raw"`
	UpdatedAt       *time.Time             `json:"updated_at,omitempty" bson:"updated_at"`
}

type GetTokenListRes struct {
	Pagination *elastic.Pagination `json:"pagination"`
	Data       *[]Token            `json:"data,omitempty"`
}

type TokeReq struct {
	IdToken      string
	ContractAddr string
}
