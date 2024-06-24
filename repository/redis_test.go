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

	t.Run("happy_test", func(t *testing.T) {

	})

	t.Run("fail_to_insert", func(t *testing.T) {

	})

}

func TestRedisGet(t *testing.T) {
	t.Run("happy_test", func(t *testing.T) {

	})

	t.Run("fail_to_get", func(t *testing.T) {

	})
}
