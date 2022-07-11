package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionRepo struct {
	*mongo.Collection
}
