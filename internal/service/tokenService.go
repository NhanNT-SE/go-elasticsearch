package service

import (
	"context"
	"log"
	"marketplace-backend/config"
	"marketplace-backend/internal/model"
	"marketplace-backend/pkg/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenSrv interface {
	GetTokenList(ctx context.Context, req *model.NFTSearchReq) (model.GetTokenListRes, error)
}

type TokenRepo struct {
	collection    *mongo.Collection
	storageNFTSrv storage.StorageNFTSrv
}

func NewTokenSrv(db *mongo.Database, cfg config.MongoConfig, storageNFTSrv storage.StorageNFTSrv) TokenSrv {
	collection := db.Client().Database(cfg.Database).Collection("token")
	return &TokenRepo{
		collection:    collection,
		storageNFTSrv: storageNFTSrv,
	}
}

func (t *TokenRepo) GetTokenList(ctx context.Context, req *model.NFTSearchReq) (model.GetTokenListRes, error) {
	var resp model.GetTokenListRes

	dataSearch, err := t.storageNFTSrv.SearchByQuery(ctx, *req)
	if err != nil {
		return resp, err
	}
	collection := t.collection
	objIds := []primitive.ObjectID{}
	for _, id := range dataSearch.Data {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return resp, err
		}
		objIds = append(objIds, objId)
	}

	filter := bson.M{"_id": bson.M{"$in": objIds}}
	m := primitive.M{}

	for _, field := range req.ResponseConfig.Source {
		m[field] = 1
	}
	opts := options.Find().SetProjection(m)

	var results []model.Token
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return resp, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		return resp, err
	}
	resp.Pagination = &dataSearch.Pagination
	resp.Data = &results

	t.getTokenId(ctx, model.TokeReq{IdToken: "20", ContractAddr: "0x441840B17D976bb6AB3C708788CE12D6b8A5CC49"})
	return resp, nil
}

// func (t *TokenRepo) CreateToken(ctx context.Context, tokenReq model.Token) {
// 	collection := t.collection

// }

func (t *TokenRepo) getTokenId(ctx context.Context, tokenReq model.TokeReq) (primitive.ObjectID, error) {
	collection := t.collection
	filter := bson.M{"id": tokenReq.IdToken, "contract_address": tokenReq.ContractAddr}

	var result bson.M
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		log.Println("getTokenId", err)
		return primitive.ObjectID{}, err
	}
	id := result["_id"].(primitive.ObjectID)
	return id, nil
}
