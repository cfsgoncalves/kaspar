package api

import (
	api "kaspar/api/handlers"
	"kaspar/configuration"
	"kaspar/repository"
	usecase "kaspar/usecase/implementation"

	"github.com/gin-gonic/gin"
)

func HTTPRouteEndpoints() *gin.Engine {
	cache := repository.NewRedis()
	stockApi := api.NewStockApi(usecase.NewStockRedditApi(cache))
	system := usecase.NewSystemMonitoring(cache)

	gin.SetMode(configuration.GetEnvAsString("GIN_MODE", "debug"))

	router := gin.Default()
	//Get a specific stock with a recomendation
	router.GET("/v1/stocks/:name/*date", stockApi.GetStockByNameAndOptionalDate)
	router.GET("/health", system.Health)
	router.GET("/ping", system.Ping)

	return router
}
