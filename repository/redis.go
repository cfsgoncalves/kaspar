package repository

import (
	"context"
	"kaspar/configuration"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Redis struct {
	redis redis.Client
}

func NewRedis() *Redis {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     configuration.GetEnvAsString("REDIS_ADDRESS", "localhost:6379"), // use default Addr
		Password: configuration.GetEnvAsString("REDIS_PASSWORD", ""),              // no password set
		DB:       configuration.GetEnvAsInt("DB", 0),                              // use default DB
	})

	return &Redis{redis: *redisClient}
}

func (r *Redis) Insert(key string, value string) error {
	TTL := time.Duration(configuration.GetEnvAsInt("REDIS_TTL", 15*60)) * time.Second

	err := r.redis.Set(ctx, key, value, TTL).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(key string) (string, error) {
	val, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, err

}

func (r *Redis) Ping() bool {

	_, err := r.redis.Ping(context.Background()).Result()

	if err != nil {
		//log error
		return false
	}

	return true

}
