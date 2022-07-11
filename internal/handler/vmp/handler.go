package vmp

import (
	"context"

	"marketplace-backend/internal/repo"
	"marketplace-backend/model"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"

	"marketplace-backend/config"
	"marketplace-backend/internal/must"
	"marketplace-backend/pkg/logger"

	"github.com/elastic/go-elasticsearch/v8"
)

type Handler struct {
	cfg config.ServerConfig
	log zerolog.Logger

	db          *mongo.Database
	redisClient redis.UniversalClient
	esClient    *elasticsearch.Client

	profileRepo    *repo.ProfileRepo
	tokenRepo      *repo.TokenRepo
	categoryRepo   *repo.Repo[model.Category]
	collectionRepo *repo.Repo[model.Collection]
	listingRepo    *repo.Repo[model.Listing]
}

func NewRPCHandler(cfg config.ServerConfig) *Handler {
	ctx := context.Background()
	db := must.ConnectMongoDB(ctx, cfg.Mongo)
	redisClient := must.ConnectRedis(ctx, cfg.Redis)
	esClient := must.ConnectElasticsearch(cfg.Elasticsearch)

	return &Handler{
		cfg: cfg,
		log: logger.New(),

		db:          db,
		redisClient: redisClient,
		esClient:    esClient,

		profileRepo:    repo.NewProfileRepo(db),
		tokenRepo:      repo.NewTokenRepo(db, esClient),
		categoryRepo:   repo.NewRepo[model.Category](db),
		collectionRepo: repo.NewRepo[model.Collection](db),
		listingRepo:    repo.NewRepo[model.Listing](db),
	}
}
