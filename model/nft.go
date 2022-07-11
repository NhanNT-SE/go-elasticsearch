package model

type NFT struct {
	//general
	TokenId              uint64      `bson:"token_id" json:"tokenId"`
	TokenName            string      `bson:"token_name" json:"tokenName"`
	CollectionName       string      `bson:"collection_name" json:"collectionName"`
	ContractAddress      string      `bson:"contract_address" json:"contractAddress"`
	IsCollectionVerified bool        `bson:"is_collection_verified" json:"isCollectionVerified,omitempty"`
	SaleType             string      `bson:"sale_type" json:"saleType,omitempty"` // NOT_FOR_SALE | ON_SALE | ON_AUCTION | FOR_RENT
	TokenStandard        string      `bson:"token_standard" json:"tokenStandard"`
	Blockchain           string      `bson:"blockchain" json:"blockchain"`
	Metadata             interface{} `bson:"metadata" json:"metadata,omitempty"`

	// details
	Currency       string `bson:"currency" json:"currency,omitempty"`
	ListingEndTime uint64 `bson:"listing_end_time" json:"listingEndTime,omitempty"`
	Owner          string `bson:"owner" json:"owner,omitempty"`
	OwnerAddress   string `bson:"owner_address" json:"ownerAddress,omitempty"`
	IsFavorite     bool   `bson:"is_favorite" json:"isFavorite,omitempty"`
	NumOfFavorite  uint64 `bson:"num_of_favorite" json:"numOfFavorite,omitempty"`
	NumOfView      uint64 `bson:"num_of_view" json:"numOfView,omitempty"`
	TokenUrl       string `bson:"token_url" json:"tokenUrl,omitempty"`
	Description    string `bson:"description" json:"description,omitempty"`

	// ON_SALE
	Price string `bson:"price" json:"price,omitempty"`

	// ON_AUCTION
	StartingPrice string `bson:"starting_price" json:"startingPrice,omitempty"`
	HighestBid    string `bson:"highest_bid" json:"highestBid,omitempty"`
	YourLastBid   string `bson:"your_last_bid" json:"yourLastBid,omitempty"`
	StatingDate   uint64 `bson:"stating_date" json:"statingDate,omitempty"`
	FinalPrice    string `bson:"final_price" json:"finalPrice,omitempty"`

	// properties
	Properties map[string]Property `json:"properties"`
}
