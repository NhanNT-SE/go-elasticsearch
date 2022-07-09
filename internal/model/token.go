package model

import (
	"time"
)

type Token struct {
	ContractAddress string    `bson:"contract_address"`
	ID              string    `bson:"id"`
	Owner           string    `bson:"owner"`
	PreviousOwner   string    `bson:"previous_owner"`
	Metadata        any       `bson:"metadata"`
	MetadataRaw     string    `bson:"metadata_raw"`
	UpdatedAt       time.Time `bson:"updated_at"`
}

func (Token) CollectionName() string {
	return "token"
}
