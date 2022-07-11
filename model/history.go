package model

type ActivityNftHistory struct {
	CollectionAddress string `bson:"collection_address" json:"collectionAddress"`
	TokenID           uint64 `bson:"token_id" json:"tokenID"`
	TxHash            string `bson:"tx_hash" json:"txHash"`
	Event             string `bson:"event" json:"event"`
	Price             string `bson:"price" json:"price"`
	Currency          string `bson:"currency" json:"currency"`
	From              string `bson:"from" json:"from"`
	To                string `bson:"to" json:"to"`
	Timestamp         uint64 `bson:"timestamp" json:"timestamp"`
}

func (ActivityNftHistory) CollectionName() string {
	return ColActivityHistory
}

type BiddingNftHistory struct {
	CollectionAddress string `bson:"collection_address" json:"collectionAddress"`
	TokenID           uint64 `bson:"token_id" json:"tokenID"`
	BidPrice          string `bson:"bid_price" json:"bidPrice"`
	USDPrice          string `bson:"usd_price" json:"usdPrice"`
	FloorPrice        string `bson:"floor_price" json:"floorPrice"`
	Date              uint64 `bson:"date" json:"expiration"`
	From              string `bson:"from" json:"from"`
	FromAddress       string `bson:"from_address" json:"fromAddress"`
}

func (BiddingNftHistory) CollectionName() string {
	return ColBiddingHistory
}

type ListingNftHistory struct {
	CollectionAddress string `bson:"collection_address" json:"collectionAddress"`
	TokenID           uint64 `bson:"token_id" json:"tokenID"`
	UnitPrice         string `bson:"unit_price" json:"unitPrice"`
	USDUnitPrice      string `bson:"usd_unit_price" json:"usdUnitPrice"`
	Quantity          uint   `bson:"quantity" json:"quantity"`
	Expiration        uint64 `bson:"expiration" json:"expiration"`
	From              string `bson:"from" json:"from"`
	FromAddress       string `bson:"from_address" json:"fromAddress"`
}

func (ListingNftHistory) CollectionName() string {
	return ColListingHistory
}

type OfferNftHistory struct {
	CollectionAddress string `bson:"collection_address" json:"collectionAddress"`
	TokenID           uint64 `bson:"token_id" json:"tokenID"`
	Price             string `bson:"price" json:"price"`
	USDPrice          string `bson:"usd_price" json:"usdPrice"`
	FloorDifference   string `bson:"floor_difference" json:"floorDifference"`
	Expiration        uint64 `bson:"expiration" json:"expiration"`
	From              string `bson:"from" json:"from"`
	FromAddress       string `bson:"from_address" json:"fromAddress"`
}

func (OfferNftHistory) CollectionName() string {
	return ColOfferHistory
}
