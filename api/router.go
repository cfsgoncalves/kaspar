package api

import (
	"kaspar/repository"
	usecase "kaspar/usecase/implementation"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func NewRouter() {
	cache := repository.NewRedis()

	stocks := usecase.NewStocksImplementation(cache)
	system := usecase.NewSystemImplementation(cache)

	router := gin.Default()
	//Recommend all stocks for a specific date
	router.GET("/v1/stocks/*date", stocks.GetStocks)
	//Recommend the best stocks to buy for a specific date
	router.GET("/v1/recommend/buy/*date", nil)
	//Recommend the best stocks to sell for a specific date
	router.GET("/v1/recommend/sell/*date", nil)
	router.GET("/health", system.Health)
	router.GET("/ping", system.Ping)

	router.Run("localhost:8080")
}
