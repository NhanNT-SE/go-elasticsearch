package repo

import (
	"context"
	"marketplace-backend/model"
	"marketplace-backend/pkg/logger"
	"marketplace-backend/pkg/storage"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenRepo struct {
	log           zerolog.Logger
	repo          *Repo[model.Token]
	storageNFTSrv storage.StorageNFTSrv
}

func NewTokenRepo(db *mongo.Database, es *elasticsearch.Client) *TokenRepo {
	storageNFTSrv := storage.NewStorageNFTSrv(es, time.Second*10)
	return &TokenRepo{
		log:           logger.New(),
		repo:          NewRepo[model.Token](db),
		storageNFTSrv: storageNFTSrv,
	}
}

func (r *TokenRepo) GetTokenList(ctx context.Context, req *model.NFTIndexSearchReq) (model.NFTIndexSearchRes, error) {
	var resp model.NFTIndexSearchRes

	dataSearch, err := r.storageNFTSrv.SearchByQuery(ctx, *req)
	if err != nil {
		return resp, err
	}
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

	results, err := r.repo.FindAll(ctx, filter, opts)
	if err != nil {
		return resp, err
	}

	resp.Pagination = &dataSearch.Pagination
	resp.Data = &results

	return resp, nil
}

func (r *TokenRepo) UpdateToken(ctx context.Context, token *model.Token) error {
	filter := bson.M{
		model.FTokenContractAddress: token.ContractAddress,
		model.FTokenID:              token.TokenID,
		model.FTokenOwner:           token.Owner,
	}
	now := time.Now()
	token.UpdatedAt = &now
	result, err := r.repo.UpdateOne(ctx, filter, bson.M{"$set": token}, &options.UpdateOptions{Upsert: &TRUE})
	if err != nil {
		return err
	}

	if result.UpsertedID == nil {
		tokenUpdate, err := r.repo.FindOne(ctx, filter)
		if err != nil {
			return err
		}
		token.ID = tokenUpdate.ID
		err = r.updateNFTIndex(ctx, "update", token)
		if err != nil {
			return err
		}
		return nil
	}

	token.ID = result.UpsertedID.(primitive.ObjectID)
	err = r.updateNFTIndex(ctx, "insert", token)

	if err != nil {
		return err
	}
	return nil
}

func (r *TokenRepo) DeleteToken(ctx context.Context, token *model.Token) error {
	filter := bson.M{
		model.FTokenContractAddress: token.ContractAddress,
		model.FTokenID:              token.TokenID,
		model.FTokenOwner:           token.Owner,
	}
	// Delete in mongo db
	_, err := r.repo.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Delete in elastic search
	err = r.storageNFTSrv.DeleteNFT(ctx, *token)
	if err != nil {
		return err
	}
	return nil
}

func (r *TokenRepo) updateNFTIndex(ctx context.Context, action string, token *model.Token) error {
	metadata := token.Metadata
	now := time.Now()

	nft := model.NFTIndex{
		TokenId:         token.TokenID,
		ContractAddress: token.ContractAddress,
		Owner:           token.Owner,
		SaleType:        string(token.SaleType),

		Attributes: metadata["attributes"],
		UpdatedAt:  token.UpdatedAt,
	}
	name, ok := metadata["name"].(string)
	if ok {
		nft.Name = name
	}
	description, ok := metadata["description"].(string)
	if ok {
		nft.Description = description
	}

	if action == "insert" {
		price := make(map[string]interface{})
		price["eth"] = 0
		price["vngt"] = 0
		nft.Price = price
		nft.CreatedTime = &now
		err := r.storageNFTSrv.InsertNFT(ctx, nft, token.ID.Hex())
		if err != nil {
			return err
		}
		return nil
	}

	err := r.storageNFTSrv.UpdateNFT(ctx, nft, token.ID.Hex())
	if err != nil {
		return err
	}

	return nil
}
