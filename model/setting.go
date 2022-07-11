package model

type NotificationSetting struct {
	WalletAddress     string `bson:"wallet_address" json:"walletAddress"`
	ItemSold          bool   `bson:"item_sold" json:"itemSold"`
	BidActivity       bool   `bson:"bid_activity" json:"bidActivity"`
	PriceChange       bool   `bson:"price_change" json:"priceChange"`
	AuctionExpiration bool   `bson:"auction_expiration" json:"auctionExpiration"`
	Outbid            bool   `bson:"outbid" json:"outbid"`
	SuccessPurchase   bool   `bson:"success_purchase" json:"successPurchase"`
	Newsletter        bool   `bson:"newsletter" json:"newsletter"`
}
