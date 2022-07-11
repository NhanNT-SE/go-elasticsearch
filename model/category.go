package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Description     string             `bson:"description" json:"description"`
	MainImage       string             `bson:"main_image" json:"mainImage"`
	SubImages       []string           `bson:"sub_images" json:"subImages"`
	CollectionCount uint               `bson:"collection_count" json:"collectionCount"`
}

func (c Category) CollectionName() string {
	return "category"
}
