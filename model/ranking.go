package model

type Ranking struct {
	Rank              uint   `bson:"rank" json:"rank"`
	CollectionName    string `bson:"collection_name" json:"collectionName"`
	CollectionAddress string `bson:"collection_address" json:"collectionAddress"`
	CreatorName       string `bson:"creator_name" json:"creatorName"`
	IsCreatorVerified bool   `bson:"is_creator_verified" json:"isCreatorVerified"`
	CollectionUrl     string `bson:"collection_url" json:"collectionUrl"`
	Volume            string `bson:"volume" json:"volume"`
	Currency          string `bson:"currency" json:"currency"`
	NumOfSale         string `bson:"num_of_sale" json:"numOfSale"`
	NumOfItem         string `bson:"num_of_item" json:"numOfItem"`
	FloorPrice        string `bson:"floor_price" json:"floorPrice"`
}
