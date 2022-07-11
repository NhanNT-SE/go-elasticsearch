package model

import (
	"encoding/json"
	"os"
)

type Marketplace struct {
	ContractAddress       string `bson:"contract_address"`
	LastSyncedBlockNumber int64  `bson:"last_synced_block_number"`
}

func (m Marketplace) CollectionName() string {
	return "marketplace" + os.Getenv("COLLECTION_NAME_SUFFIX")
}

type MarketplaceEventName string

const (
	MarketplaceEventPaymentTokenUpdated = "PaymentTokenUpdated"
	MarketplaceEventSupportedNFTUpdated = "SupportedNFTUpdated"
	MarketplaceEventRoleGranted         = "RoleGranted"
	MarketplaceEventItemListed          = "ItemListed"
	MarketplaceEventItemBought          = "ItemBought"
	MarketplaceEventAuctionCreated      = "AuctionCreated"
	MarketplaceEventAuctionCanceled     = "AuctionCanceled"
	MarketplaceEventAuctionClaimed      = "AuctionClaimed"
	MarketplaceEventAuctionItemSold     = "AuctionItemSold"
	MarketplaceEventOfferCreated        = "OfferCreated"
	MarketplaceEventOfferCanceled       = "OfferCanceled"
	MarketplaceEventOfferAccepted       = "OfferAccepted"
	MarketplaceEventExpiredOfferClaimed = "ExpiredOfferClaimed"
	MarketplaceEventBidPlaced           = "BidPlaced"
	MarketplaceEventListingCanceled     = "ListingCanceled"
	MarketplaceEventExpiredItemClaimed  = "ExpiredItemClaimed"
)

type MarketplaceEvent struct {
	Name            string `json:"name"`
	ContractAddress string `json:"contractAddress"`
	BlockNumber     uint64 `json:"blockNumber"`
	LogIndex        uint   `json:"logIndex"`
}

func (e *MarketplaceEvent) Bytes() []byte {
	if e == nil {
		return nil
	}
	b, _ := json.Marshal(e)
	return b
}
