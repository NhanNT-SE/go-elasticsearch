package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListingStatus string

const (
	ListingStatusOpen     ListingStatus = "OPEN"
	ListingStatusCanceled ListingStatus = "CANCELED"
	ListingStatusClosed   ListingStatus = "CLOSED"
)

type Listing struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MarketID        string             `bson:"market_id" json:"marketId"`
	ContractAddress string             `bson:"contract_address" json:"contractAddress"`
	TokenID         string             `bson:"token_id" json:"tokenId"`
	SellerAddress   string             `bson:"seller_address" json:"sellerAddress"`
	PaymentToken    string             `bson:"payment_token" json:"paymentToken"`
	Price           string             `bson:"price" json:"price"`
	Quantity        string             `bson:"quantity" json:"quantity"`
	Status          ListingStatus      `bson:"status" json:"status"`
	ExpiredAt       time.Time          `bson:"expired_at" json:"expiredAt"`
	Expiration      time.Duration      `bson:"expiration" json:"expiration"`
	TxHash          string             `bson:"tx_hash" json:"txHash"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updatedAt"`
}

func (Listing) CollectionName() string {
	return "listing"
}

// type ListingHistory struct {
// 	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
// 	ContractAddress string             `bson:"contract_address" json:"contractAddress"`
// 	TokenID         string             `bson:"token_id" json:"tokenId"`
// 	PaymentToken    string             `bson:"payment_token" json:"paymentToken"`
// 	Price           string             `bson:"price" json:"price"`
// 	Quantity        string             `bson:"quantity" json:"quantity"`
// 	ExpiredAt       time.Time          `bson:"expired_at" json:"expiredAt"`
// 	Expiration      time.Duration      `bson:"expiration" json:"expiration"`
// 	UpdatedAt       time.Time          `bson:"updated_at" json:"updatedAt"`
// }
//
// func (ListingHistory) CollectionName() string {
// 	return "listing_history"
// }
