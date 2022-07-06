package vmp

import (
	"marketplace-backend/config"
	"marketplace-backend/internal/must"
	"marketplace-backend/pkg/logger"
	"marketplace-backend/pkg/storage"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog"
)

type Handler struct {
	cfg config.ServerConfig
	log zerolog.Logger

	// db *mongo.Database
	// redisClient redis.UniversalClient
	esClient      *elasticsearch.Client
	storageNFTSrv storage.StorageNFTSrv
}

func NewRPCHandler(cfg config.ServerConfig) *Handler {
	// ctx := context.Background()
	// db := must.ConnectMongoDB(ctx, cfg.Mongo)
	// redisClient := must.ConnectRedis(ctx, cfg.Redis)
	esClient := must.ConnectElasticsearch(&cfg.Elasticsearch)
	storageNFTSrv := storage.NewStorageNFTSrv(esClient, time.Second*10)
	return &Handler{
		cfg: cfg,
		log: logger.New(),

		// db:          db,
		// redisClient: redisClient,
		esClient:      esClient,
		storageNFTSrv: storageNFTSrv,
	}
}
