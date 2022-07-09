package vmp

import (
	"context"
	"marketplace-backend/config"
	"marketplace-backend/internal/must"
	"marketplace-backend/internal/service"
	"marketplace-backend/pkg/logger"
	"marketplace-backend/pkg/storage"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	cfg config.ServerConfig
	log zerolog.Logger

	db *mongo.Database
	// redisClient redis.UniversalClient
	esClient      *elasticsearch.Client
	tokenSrv      service.TokenSrv
	storageNFTSrv storage.StorageNFTSrv
}

func NewRPCHandler(cfg config.ServerConfig) *Handler {
	ctx := context.Background()
	db := must.ConnectMongoDB(ctx, cfg.Mongo)
	// redisClient := must.ConnectRedis(ctx, cfg.Redis)
	esClient := must.ConnectElasticsearch(&cfg.Elasticsearch)
	storageNFTSrv := storage.NewStorageNFTSrv(esClient, time.Second*10)
	tokenSrv := service.NewTokenSrv(db, cfg.Mongo, storageNFTSrv)
	return &Handler{
		cfg: cfg,
		log: logger.New(),

		db: db,
		// redisClient: redisClient,
		tokenSrv:      tokenSrv,
		esClient:      esClient,
		storageNFTSrv: storageNFTSrv,
	}
}
