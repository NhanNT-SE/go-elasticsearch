package model

import (
	"math/big"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaleType string

const (
	SaleTypeNotForSale SaleType = "NOT_FOR_SALE"
	SaleTypeOnSale     SaleType = "ON_SALE"
	SaleTypeOnAuction  SaleType = "ON_AUCTION"
	SaleTypeForRent    SaleType = "FOR_RENT"
)

type Token struct {
	ID              primitive.ObjectID     `json:"-" bson:"_id,omitempty"`
	TokenID         string                 `json:"token_id,omitempty" bson:"token_id"`
	ContractAddress string                 `json:"contract_address,omitempty" bson:"contract_address"`
	Owner           string                 `json:"owner,omitempty" bson:"owner"`
	PreviousOwner   string                 `json:"previousOwner,omitempty" bson:"previous_owner"`
	Value           string                 `json:"value,omitempty" bson:"value"`
	Type            TokenStandard          `json:"type,omitempty" bson:"type"`
	Metadata        map[string]interface{} `json:"metadata,omitempty" bson:"metadata"`
	MetadataRaw     string                 `json:"metadata_raw,omitempty" bson:"metadata_raw"`
	SaleType        SaleType               `json:"sale_type,omitempty" bson:"sale_type"`
	UpdatedAt       *time.Time             `json:"updatedAt,omitempty" bson:"updated_at"`
}

func (Token) CollectionName() string {
	return "token" + os.Getenv("COLLECTION_NAME_SUFFIX")
}

func (t *Token) IDAsBigInt() *big.Int {
	v, _ := big.NewInt(0).SetString(t.TokenID, 10)
	return v
}

func (t *Token) ValueAsBigInt() *big.Int {
	v, _ := big.NewInt(0).SetString(t.Value, 10)
	return v
}
