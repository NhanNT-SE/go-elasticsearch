package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BiddingStatus string

const (
	BiddingStatusOpen     BiddingStatus = "OPEN"
	BiddingStatusCanceled BiddingStatus = "CANCELED"
	BiddingStatusExpired  BiddingStatus = "EXPIRED"
	BiddingStatusOutBid   BiddingStatus = "OUT_BID"
)

type Bidding struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ContractAddress     string             `bson:"contract_address" json:"contractAddress"`
	TokenID             string             `bson:"token_id" json:"tokenId"`
	AuctionID           string             `bson:"auction_id" json:"auctionId"`
	PaymentTokenAddress string             `bson:"payment_token_address" json:"paymentTokenAddress"`
	BidPrice            string             `bson:"bid_price" json:"bidPrice"`
	BidderAddress       string             `bson:"bidder_address" json:"bidderAddress"`
	Status              BiddingStatus      `bson:"status" json:"status"`
	TxHash              string             `bson:"tx_hash" json:"txHash"`
	UpdatedAt           time.Time          `bson:"updated_at" json:"updatedAt"`
}

func (Bidding) CollectionName() string {
	return "bidding"
}
