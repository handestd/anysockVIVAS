package myredis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func SetSinglekey(key string, value interface{}) error {
	ctx := context.Background()
	err2 := Client.Set(ctx, key, value, 0).Err()
	if err2 != nil {
		return errors.New(err2.Error())
	}
	return nil
}
func GetSingleKey(key string) (interface{}, error) {
	ctx := context.Background()
	val, err2 := Client.Get(ctx, "foo").Result()
	if err2 != nil {
		return nil, errors.New(err2.Error())
	}
	return val, nil
}
