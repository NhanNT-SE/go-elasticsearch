package model

import "github.com/ethereum/go-ethereum/core/types"

// type RawEvent struct {
// 	// address of the contract that generated the event
// 	Address common.Address `json:"address" bson:"address"`
// 	// list of topics provided by the contract.
// 	Topics []common.Hash `json:"topics" bson:"topics"`
// 	// supplied by the contract, usually ABI-encoded
// 	Data []byte `json:"data" bson:"data"`
//
// 	// Derived fields. These fields are filled in by the node
// 	// but not secured by consensus.
// 	// block in which the transaction was included
// 	BlockNumber uint64 `json:"block_number" bson:"block_number"`
// 	// hash of the transaction
// 	TxHash common.Hash `json:"transaction_hash" bson:"transaction_hash"`
// 	// index of the transaction in the block
// 	TxIndex uint `json:"transaction_index" bson:"transaction_index"`
// 	// hash of the block in which the transaction was included
// 	BlockHash common.Hash `json:"block_hash" bson:"block_hash"`
// 	// index of the log in the block
// 	Index uint `json:"log_index" bson:"log_index"`
//
// 	Name string `json:"name" bson:"name"`
// }

type RawEvent struct {
	types.Log `bson:",inline"`
	Name      string `json:"name" bson:"name"`
}

func (r RawEvent) CollectionName() string {
	return "raw_event"
}
