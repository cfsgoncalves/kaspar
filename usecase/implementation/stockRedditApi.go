package usecase

import (
	"encoding/json"
	"errors"
	"io"
	"kaspar/configuration"
	"kaspar/entities"
	"kaspar/repository"
	"net/http"

	"github.com/rs/zerolog/log"
)

type StockRedditApi struct {
	Cache repository.Cache
}

func NewStockRedditApi(cache repository.Cache) *StockRedditApi {
	return &StockRedditApi{Cache: cache}
}

// Add bad input
func (s *StockRedditApi) GetStockByName(date string, stockName string) (any, error) {
	var stockList []entities.RedditStock

	//Check for cache if exists
	stockListCache, err := s.Cache.Get(date)
	if err != nil {
		log.Error().Msgf("usecase.GetStockByName(): Error yield acessing cache service. Error: %s", err)
	}

	// Return the values on cache with 200 and value
	if stockListCache != "" {
		err := json.Unmarshal([]byte(stockListCache), &stockList)

		if err != nil {
			log.Error().Msgf("usecase.GetStockByName(): Error yield while unmarshaling stock list. Error: %s", err)
			return nil, err
		}

		selectedStock, err := s.findStockByName(stockList, stockName)

		if err != nil {
			log.Error().Msgf("usecase.GetStockByName(): Error yield finding stock by name. Error: %s", err)
			return nil, err
		}
		return selectedStock, nil
	}

	//Get stocks from reddit api and save on cache
	stockList, err = s.fetchFromRedditApi(date, s.Cache)

	if err != nil {
		log.Error().Msgf("usecase.GetStockByName(): Error yield fething from reddit api. Error: %s", err)
		return nil, err
	}

	selectedStock, err := s.findStockByName(stockList, stockName)

	if err != nil {
		log.Error().Msgf("usecase.GetStockByName(): Error yield finding stock by name. Error: %s", err)
		return nil, err
	}

	return selectedStock, nil
}

func (s *StockRedditApi) findStockByName(stockList []entities.RedditStock, stockName string) (any, error) {

	for _, stock := range stockList {
		if stock.Ticker == stockName {
			return stock, nil
		}
	}

	log.Error().Msgf("usecase.findStockByName(): Could not find the stock that was input by the user")
	return nil, errors.New("could not find the stock that was input by the user")
}

func (s *StockRedditApi) fetchFromRedditApi(date string, cache repository.Cache) ([]entities.RedditStock, error) {
	var stockList []entities.RedditStock
	REDDIT_API_URL := configuration.GetEnvAsString("REDDIT_API_URL", "https://tradestie.com/api/v1/apps/reddit?date=") + date

	resp, err := http.Get(REDDIT_API_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error().Msgf("usecase.fetchFromRedditApi(): Error yield while reading body from GET request. Error: %s", err)
		return nil, err
	}

	// Consider putting it in parallel execution
	// Would need to create a new method on the interface to do it. Something like parallel insert.
	err = cache.Insert(date, string(body))

	if err != nil {
		log.Error().Msgf("usecase.fetchFromRedditApi(): Error yield while inserting on cache service. Error: %s", err)
	}

	err = json.Unmarshal([]byte(body), &stockList)

	if err != nil {
		log.Error().Msgf("usecase.fetchFromRedditApi(): Error yield while unmarshaling body from GET request. Error: %s", err)
		return nil, err
	}

	return stockList, err
}
