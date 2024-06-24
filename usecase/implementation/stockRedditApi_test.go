package usecase

import "testing"

func TestGetStockByName(t *testing.T) {
	t.Run("happy_testing_not_cache", func(t *testing.T) {

	})

	t.Run("happy_testing_cache_hit", func(t *testing.T) {

	})

	t.Run("could_not_get_cache", func(t *testing.T) {

	})

	t.Run("fail_to_find_stock_but_still_cached", func(t *testing.T) {

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
