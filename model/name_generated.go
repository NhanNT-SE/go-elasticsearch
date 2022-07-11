// Code generated. DO NOT EDIT.

package model

const (
	FObjectID = "_id"
)

const (
	ColCollection = "collection"

	FCollectionName                  = "name"
	FCollectionContractAddress       = "contract_address"
	FCollectionCreatorName           = "creator_name"
	FCollectionCreatorAddress        = "creator_address"
	FCollectionTokenStandard         = "token_standard"
	FCollectionLastSyncedBlockNumber = "last_synced_block_number"
	FCollectionUpdatedAt             = "updated_at"
)

const (
	ColCategory = "category"

	FCategoryID          = "_id"
	FCategoryName        = "name"
	FCategoryDescription = "description"
)

const (
	ColMarketplace = "marketplace"

	FMarketplaceContractAddress       = "contract_address"
	FMarketplaceLastSyncedBlockNumber = "last_synced_block_number"
)

const (
	ColRawEvent = "raw_event"

	FRawEventName        = "name"
	FRawEventAddress     = "address"
	FRawEventTopics      = "topics"
	FRawEventData        = "data"
	FRawEventBlockNumber = "blocknumber"
	FRawEventTxHash      = "txhash"
	FRawEventTxIndex     = "txindex"
	FRawEventBlockHash   = "blockhash"
	FRawEventIndex       = "index"
	FRawEventRemoved     = "removed"
)

const (
	ColProfile = "profile"

	FUsername      = "username"
	FWalletAddress = "wallet_address"
)

const (
	ColRanking         = "ranking"
	ColActivityHistory = "activity_history"
	ColListingHistory  = "listing_history"
	ColOfferHistory    = "offer_history"
	ColBiddingHistory  = "bidding_history"

	FCollectionAddress = "collection_address"
	FTxHash            = "tx_hash"
	FEvent             = "event"
)

const (
	ColConfig = "config"

	FConfigKey  = "key"
	FConfigData = "data"
)

const (
	ColToken = "token"

	FTokenID              = "id"
	FTokenContractAddress = "contract_address"
	FTokenTokenID         = "token_id"
	FTokenOwner           = "owner"
	FTokenPreviousOwner   = "previous_owner"
	FTokenUpdatedAt       = "updated_at"
	FMetadata             = "metadata"
	FMetadataRaw          = "metadata_raw"
	FTokenValue           = "value"
	FTokenType            = "type"
	FTokenSaleType        = "sale_type"
)

const (
	FPaymentTokenContractAddress = "contract_address"
	FPaymentTokenAllowed         = "allowed"
	FPaymentTokenUpdatedAt       = "updated_at"
)

const (
	FListingMarketID        = "market_id"
	FListingContractAddress = "contract_address"
	FListingTokenID         = "token_id"
	FListingStatus          = "status"
	FListingUpdatedAt       = "updated_at"
	FListingTxHash          = "tx_hash"
	FListingQuantity        = "quantity"
)

const (
	FOfferMarketID        = "market_id"
	FOfferContractAddress = "contract_address"
	FOfferTokenID         = "token_id"
	FOfferStatus          = "status"
	FOfferUpdatedAt       = "updated_at"
	FOfferTxHash          = "tx_hash"
)

const (
	FBiddingAuctionID       = "auction_id"
	FBiddingContractAddress = "contract_address"
	FBiddingTokenID         = "token_id"
	FBiddingStatus          = "status"
	FBiddingBidderAddress   = "bidder_address"
	FBiddingUpdatedAt       = "updated_at"
	FBiddingTxHash          = "tx_hash"
)

const (
	FAuctionMarketID        = "market_id"
	FAuctionContractAddress = "contract_address"
	FAuctionTokenID         = "token_id"
	FAuctionStatus          = "status"
	FAuctionUpdatedAt       = "updated_at"
	FAuctionTxHash          = "tx_hash"
)
