package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model interface {
	CollectionName() string
}

type Repo[T Model] struct {
	*mongo.Collection
}

func NewRepo[T Model](db *mongo.Database) *Repo[T] {
	var t T
	return &Repo[T]{Collection: db.Collection(t.CollectionName())}
}

func (r *Repo[T]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*T, error) {
	var m T
	err := r.Collection.FindOne(ctx, filter, opts...).Decode(&m)
	return &m, err
}

func (r *Repo[T]) FindAll(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	cs, err := r.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	ms := make([]T, 0)
	err = cs.All(ctx, &ms)
	return ms, err
}
