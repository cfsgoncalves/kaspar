package usecase

import "github.com/gin-gonic/gin"

type Stocks interface {
	GetStocks(*gin.Context)
	GetStocksToBuy(*gin.Context)
	GetStocksToSell(*gin.Context)
}
