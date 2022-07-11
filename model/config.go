package model

import "time"

type Config struct {
	Key       string    `bson:"key"`
	Data      []byte    `bson:"data"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (Config) CollectionName() string {
	return "config"
}
