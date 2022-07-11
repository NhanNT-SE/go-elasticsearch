package model

import (
	"os"
	"time"
)

type TokenStandard string

const (
	TokenStandardUnknown = "UNKNOWN"
	TokenStandardERC721  = "ERC721"
	TokenStandardERC1155 = "ERC1155"
)

var mapNFTTypeToTokenStandard = map[uint8]TokenStandard{
	1: TokenStandardERC721,
	2: TokenStandardERC1155,
}

func TokenStandardFromNFTType(nftType uint8) TokenStandard {
	if v, ok := mapNFTTypeToTokenStandard[nftType]; ok {
		return v
	}
	return TokenStandardUnknown
}

type Collection struct {
	Name                  string        `bson:"name" json:"name"`
	ContractAddress       string        `bson:"contract_address" json:"contractAddress"`
	TokenName             string        `bson:"token_name" json:"tokenName"`
	TokenSymbol           string        `bson:"token_symbol" json:"tokenSymbol"`
	TokenStandard         TokenStandard `bson:"token_standard" json:"tokenStandard"`
	LastSyncedBlockNumber int64         `bson:"last_synced_block_number" json:"-"`
	UpdatedAt             time.Time     `bson:"updated_at" json:"updatedAt"`

	// creator
	CreatorName        string `bson:"creator_name" json:"creatorName"`
	CreatorAddress     string `bson:"creator_address" json:"creatorAddress"`
	CreatorDescription string `bson:"creator_description" json:"creatorDescription"`
	CreatorVerified    bool   `bson:"creator_verified" json:"creatorVerified"`

	// general
	CollectionUrl     string `bson:"collection_url" json:"collectionUrl,omitempty"`
	CollectionBanner  string `bson:"collection_banner" json:"collectionBanner,omitempty"`
	CollectionDetails string `bson:"collection_details" json:"collectionDetails,omitempty"`

	// Analyst
	HighestSale string `bson:"highest_sale" json:"highestSale,omitempty"`
	FloorPrice  string `bson:"floor_price" json:"floorPrice,omitempty"`
	Volume      string `bson:"volume" json:"volume,omitempty"`
	MarketCap   string `bson:"market_cap" json:"marketCap,omitempty"`
	Currency    string `bson:"currency" json:"currency,omitempty"`
	NumOfViews  uint64 `bson:"num_of_views" json:"numOfViews,omitempty"`
	NumOfOwners uint64 `bson:"num_of_owners" json:"numOfOwners,omitempty"`
	NumOfItems  uint64 `bson:"num_of_items" json:"numOfItems,omitempty"`

	// properties
	Properties map[string][]Property `bson:"properties" json:"properties,omitempty"`
}

func (c Collection) CollectionName() string {
	return "collection" + os.Getenv("COLLECTION_NAME_SUFFIX")
}
