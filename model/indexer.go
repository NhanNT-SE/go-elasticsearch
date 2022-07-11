package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	KeyToken      = "token"
	KeyCollection = "collection"
)

type IndexerTokenEvent struct {
	ID primitive.ObjectID `json:"id"`
}

type IndexerToken struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
