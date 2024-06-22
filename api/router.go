package api

import (
	"kaspar/configuration"
	"kaspar/repository"
	usecase "kaspar/usecase/implementation"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func NewRouter() {
	cache := repository.NewRedis()
	stockApi := NewStockApi(usecase.NewStockRedditApi(cache))
	system := usecase.NewSystemMonitoring(cache)

	gin.SetMode(configuration.GetEnvAsString("GIN_MODE", "debug"))

	router := gin.Default()
	//Get a specific stock with a recomendation
	router.GET("/v1/stocks/:name/*date", stockApi.GetStockByNameAndOptionalDate)
	router.GET("/health", system.Health)
	router.GET("/ping", system.Ping)

	router.Run("localhost:8080")
}
