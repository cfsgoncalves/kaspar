package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kaspar/entities"
	"kaspar/repository"
	"net/http"
)

type StockRedditApi struct {
	Cache repository.Cache
}

func NewStockRedditApi(cache repository.Cache) *StockRedditApi {
	return &StockRedditApi{Cache: cache}
}

// Incomplete
func (s *StockRedditApi) GetStockByName(date string, stockName string) (string, error) {
	var stockList []entities.RedditStock

	//Check for cache if exists
	stockListCache, err := s.Cache.Get(date)
	if err != nil {
		//Log error if exists, but don't terminate
		return "", err
	}

	// Return the values on cache with 200 and value
	if stockListCache != "" {
		err := json.Unmarshal([]byte(stockListCache), &stockList)

		if err != nil {
			return "", err
		}

		jsonStock, err := FindStockByName(stockList, stockName)

		if err != nil {
			return "", err
		}

		return jsonStock, nil
	}

	//Get stocks from reddit api and save on cache
	stockList, err = FetchFromRedditApi(date, s.Cache)

	if err != nil {
		return "", err
	}

	jsonStock, err := FindStockByName(stockList, stockName)

	if err != nil {
		return "", nil
	}
	return jsonStock, nil
}

func FindStockByName(stockList []entities.RedditStock, stockName string) (string, error) {

	for _, stock := range stockList {
		if stock.Ticker == stockName {
			value, err := json.Marshal(stock)
			if err != nil {
				return "", err
			}
			return string(value), nil
		}
	}

	return "", errors.New("could not find the stock that was input by the user")
}

func FetchFromRedditApi(date string, cache repository.Cache) ([]entities.RedditStock, error) {
	var stockList []entities.RedditStock
	REQUEST_URL := "https://tradestie.com/api/v1/apps/reddit?date=" + date

	resp, err := http.Get(REQUEST_URL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		// log could not read body
		fmt.Printf("Could not read body")
		return nil, err
	}

	//Save on Cache - Parallel - Independent from finding the stock
	go cache.Insert(date, string(body))

	err = json.Unmarshal([]byte(body), &stockList)

	if err != nil {
		return nil, err
	}

	return stockList, err
}
