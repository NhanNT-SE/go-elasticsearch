package service

import (
	"log"
	"marketplace-backend/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type TokenSrv interface {
	GetTokenList(iDList []string)
}

type TokenRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTokenSrv(db *mongo.Database, cfg config.MongoConfig) TokenSrv {
	collection := db.Client().Database(cfg.Database).Collection("token")
	return &TokenRepo{
		db:         db,
		collection: collection,
	}
}

func (t *TokenRepo) GetTokenList(iDList []string) {
	// collection := t.collection
	// // collection.Find({})
	log.Println(iDList[0])
}
