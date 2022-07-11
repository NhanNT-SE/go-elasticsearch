package repo

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"marketplace-backend/model"
	"marketplace-backend/pkg/logger"
	"reflect"
)

var (
	TRUE = true
)

type ProfileRepo struct {
	log  zerolog.Logger
	repo *Repo[model.Profile]
}

func NewProfileRepo(db *mongo.Database) *ProfileRepo {
	return &ProfileRepo{
		log:  logger.New(),
		repo: NewRepo[model.Profile](db),
	}
}

func (r *ProfileRepo) GetProfile(ctx context.Context, key string, result interface{}) error {
	if reflect.TypeOf(result).Kind() != reflect.Pointer {
		return errors.New("result must be a pointer")
	}
	if result == nil {
		return errors.New("result must be not nil")
	}

	rs, err := r.repo.FindOne(ctx, bson.M{model.FWalletAddress: key})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			r.log.Err(err).Msg("DB error")
		}
		return err
	}
	result = &rs

	return nil
}

func (r *ProfileRepo) UpdateProfile(ctx context.Context, data model.Profile) error {

	filter := bson.M{model.FWalletAddress: data.WalletAddress}
	_, err := r.repo.UpdateOne(ctx, filter, bson.M{"$set": data}, &options.UpdateOptions{Upsert: &TRUE})

	if err != nil {
		return err
	}

	return nil
}
