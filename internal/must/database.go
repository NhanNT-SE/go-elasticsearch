package must

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"marketplace-backend/config"
	"marketplace-backend/pkg/logger"
)

var (
	log = logger.New()
)

func ConnectMongoDB(ctx context.Context, cfg config.MongoConfig) *mongo.Database {
	ctx, cc := context.WithTimeout(ctx, 5*time.Second)
	defer cc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		log.Fatal().Err(err).Msg("connect mongo failed")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal().Err(err).Msg("ping mongo failed")
	}

	log.Debug().Msg("connect mongodb successfully")
	return client.Database(cfg.Database)
}

func ConnectRedis(ctx context.Context, cfg config.RedisConfig) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		DB:       cfg.DB,
		Password: cfg.Password,
		Username: cfg.User,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("ping redis failed")
	}

	log.Debug().Msg("connect redis successfully")
	return redisClient
}
