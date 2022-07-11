package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OfferStatus string

const (
	OfferStatusOpen     OfferStatus = "OPEN"
	OfferStatusCanceled OfferStatus = "CANCELED"
	OfferStatusAccepted OfferStatus = "ACCEPTED"
	OfferStatusExpired  OfferStatus = "EXPIRED"
	OfferStatusClaimed  OfferStatus = "CLAIMED"
)

type Offer struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ContractAddress     string             `bson:"contract_address" json:"contractAddress"`
	TokenID             string             `bson:"token_id" json:"tokenId"`
	MarketID            string             `bson:"market_id" json:"marketId"`
	PaymentTokenAddress string             `bson:"payment_token_address" json:"paymentTokenAddress"`
	OfferPrice          string             `bson:"offer_price" json:"offerPrice"`
	OfferorAddress      string             `bson:"offeror_address" json:"offerorAddress"`
	Status              OfferStatus        `bson:"status" json:"status"`
	EndTime             time.Time          `bson:"end_time" json:"endTime"`
	TxHash              string             `bson:"tx_hash" json:"txHash"`
	UpdatedAt           time.Time          `bson:"updated_at" json:"updatedAt"`
}

func (Offer) CollectionName() string {
	return "offer"
}
