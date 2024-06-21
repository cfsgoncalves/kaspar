package repository

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	redis redis.Client
}

func NewRedis() *Redis {
	REDIS_ADDRESS := "localhost:6379"
	REDIS_PASSWORD := ""
	DB := 0

	redisClient := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDRESS,  // use default Addr
		Password: REDIS_PASSWORD, // no password set
		DB:       DB,             // use default DB
	})

	return &Redis{redis: *redisClient}
}

// https://gist.github.com/miguelmota/c4ff27419e53492a66a171cbe02fe033
func (r *Redis) Insert(key string, value string) error {
	fmt.Print("Inserting data into Redis")
	return nil
}

func (r *Redis) Get(date string) (string, error) {
	fmt.Print("Getting data from Redis")
	return "", nil
}

func (r *Redis) Ping() bool {

	_, err := r.redis.Ping().Result()

	if err != nil {
		//log error
		return false
	}

	return true

}
