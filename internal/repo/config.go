package repo

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"marketplace-backend/model"
)

type ConfigRepo struct {
	repo *Repo[model.Config]
}

func NewConfigRepo(db *mongo.Database) *ConfigRepo {
	return &ConfigRepo{
		repo: NewRepo[model.Config](db),
	}
}

func (r *ConfigRepo) GetConfigData(ctx context.Context, key string, result interface{}) error {
	if reflect.TypeOf(result).Kind() == reflect.Pointer {
		return errors.New("result must be a pointer")
	}
	if result == nil {
		return errors.New("result must be not nil")
	}

	cfg, err := r.repo.FindOne(ctx, bson.M{model.FConfigKey: key})
	if err != nil {
		return err
	}

	return json.Unmarshal(cfg.Data, result)
}
