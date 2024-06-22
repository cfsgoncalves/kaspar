package repository

import (
	"testing"

	"github.com/go-redis/redismock/v9"
)

func TestRedisInsert(t *testing.T) {
	t.Run("test_ttl", func(t *testing.T) {
		redis := NewRedis()

		db, _ := redismock.NewClientMock()

		redis.redis = *db

	})

}
