package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentToken struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	ContractAddress string             `bson:"contract_address" json:"contractAddress"`
	Allowed         bool               `bson:"allowed" json:"allowed"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updatedAt"`
}

func (PaymentToken) CollectionName() string {
	return "payment_token"
}
