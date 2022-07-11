package model

import "time"

type AuctionStatus string

const (
	AuctionStatusOpen     AuctionStatus = "OPEN"
	AuctionStatusCanceled AuctionStatus = "CANCELED"
	AuctionStatusSold     AuctionStatus = "SOLD"
	AuctionStatusExpired  AuctionStatus = "EXPIRED"
	AuctionStatusClaimed  AuctionStatus = "CLAIMED"
)

type Auction struct {
	ID                  string        `bson:"_id,omitempty" json:"id,omitempty"`
	MarketID            string        `bson:"market_id" json:"marketId"`
	ContractAddress     string        `bson:"contract_address" json:"contractAddress"`
	TokenID             string        `bson:"token_id" json:"tokenId"`
	OwnerAddress        string        `bson:"owner_address" json:"ownerAddress"`
	PaymentTokenAddress string        `bson:"payment_token_address" json:"paymentTokenAddress"`
	StartingPrice       string        `bson:"starting_price" json:"startingPrice"`
	BuyNowPrice         string        `bson:"buy_now_price" json:"buyNowPrice"`
	EndTime             time.Time     `bson:"end_time" json:"endTime"`
	Status              AuctionStatus `bson:"status" json:"status"`
	TxHash              string        `bson:"tx_hash" json:"txHash"`
	UpdatedAt           time.Time     `bson:"updated_at" json:"updatedAt"`
}

func (Auction) CollectionName() string {
	return "auction"
}
