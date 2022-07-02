package vmp

import (
	"go-elasticsearch/config"
	"go-elasticsearch/internal/must"
	"go-elasticsearch/pkg/logger"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog"
)

type Handler struct {
	cfg config.ServerConfig
	log zerolog.Logger

	// db *mongo.Database
	// redisClient redis.UniversalClient
	esClient *elasticsearch.Client
}

func NewRPCHandler(cfg config.ServerConfig) *Handler {
	// ctx := context.Background()
	// db := must.ConnectMongoDB(ctx, cfg.Mongo)
	// redisClient := must.ConnectRedis(ctx, cfg.Redis)
	esClient := must.ConnectElasticsearch(&cfg.Elasticsearch)

	return &Handler{
		cfg: cfg,
		log: logger.New(),

		// db:          db,
		// redisClient: redisClient,
		esClient: esClient,
	}
}
