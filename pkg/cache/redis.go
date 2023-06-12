package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var Client *redis.Client

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetSinglekey(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	err2 := Client.Set(ctx, key, value, expiration).Err()
	if err2 != nil {
		return errors.New(err2.Error())
	}
	return nil
}
func GetSingleKey(key string) (interface{}, error) {
	ctx := context.Background()
	val, err2 := Client.Get(ctx, key).Result()
	if err2 != nil {
		return nil, errors.New(err2.Error())
	}
	return val, nil
}
func DelKey(key string, ctx context.Context) (int64, error) {
	return Client.Del(ctx, key).Result()
}

func SetMultiple(ctx context.Context, session map[string]interface{}, keyName string) {

	for k, v := range session {
		err := Client.HSet(ctx, keyName, k, v).Err()

		if err != nil {
			panic(err)
		}
	}
}
func GetMultiple(keyName string) map[string]string {
	ctx := context.Background()
	userSession := Client.HGetAll(ctx, keyName).Val()
	return userSession
}
