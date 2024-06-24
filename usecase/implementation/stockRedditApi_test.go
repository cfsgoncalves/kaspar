package usecase

import (
	"errors"
	"kaspar/entities"
	"kaspar/repository"
	"os"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestGetStockByName(t *testing.T) {

	os.Setenv("REDIS_PASSWORD", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	os.Setenv("DB", "0")
	os.Setenv("REDIS_TTL", "10")
	os.Setenv("REDIS_SERVER", "localhost")
	os.Setenv("REDIS_PORT", "6379")

	t.Run("happy_testing_not_cache", func(t *testing.T) {
		stockTicker := "TSLA"
		cache := repository.NewRedis()

		redditApi := NewStockRedditApi(cache)

		jsonValue, err := redditApi.GetStockByName("2024-06-24", stockTicker)

		assert.Nil(t, err)
		assert.Equal(t, jsonValue.(entities.RedditStock).Ticker, stockTicker)

	})

	t.Run("happy_testing_cache_hit", func(t *testing.T) {
		stockTicker := "TSLA"
		cache := repository.NewRedis()

		redditApi := NewStockRedditApi(cache)

		jsonValue, err := redditApi.GetStockByName("2024-06-24", stockTicker)

		assert.Nil(t, err)
		assert.Equal(t, jsonValue.(entities.RedditStock).Ticker, stockTicker)

		jsonValue, err = redditApi.GetStockByName("2024-06-24", stockTicker)

		assert.Nil(t, err)
		assert.Equal(t, jsonValue.(entities.RedditStock).Ticker, stockTicker)
	})

	t.Run("could_not_get_cache", func(t *testing.T) {
		stockTicker := "TSLA"

		db, mock := redismock.NewClientMock()

		cache := repository.NewRedis()

		mock.ExpectGet("2024-06-24").SetErr(errors.New("error"))
		mock.ExpectSet("2024-06-24", "fooo", 1).SetErr(errors.New("error"))

		cache.Redis = *db

		redditApi := NewStockRedditApi(cache)

		jsonValue, err := redditApi.GetStockByName("2024-06-24", stockTicker)

		assert.Nil(t, err)
		assert.Equal(t, jsonValue.(entities.RedditStock).Ticker, stockTicker)
	})

	t.Run("fail_to_find_stock_but_still_cached", func(t *testing.T) {
		stockTicker := "TSL"
		cache := repository.NewRedis()

		redditApi := NewStockRedditApi(cache)

		_, err := redditApi.GetStockByName("2024-06-24", stockTicker)

		assert.Error(t, err, "")

		cache_value, err := redditApi.Cache.Get("2024-06-24")

		assert.Nil(t, err)
		assert.NotNil(t, cache_value)
	})
}

func TestFindStockByName(t *testing.T) {
	t.Run("happy_testing", func(t *testing.T) {

	})

	t.Run("fail_to_find_stock", func(t *testing.T) {

	})
}

func TestFetchFromRedditApi(t *testing.T) {
	t.Run("happy_testing", func(t *testing.T) {

	})

	t.Run("fail_to_fetch_from_reddit", func(t *testing.T) {

	})

	t.Run("cache_not_found", func(t *testing.T) {

	})
}
