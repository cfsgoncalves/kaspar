package usecase

import (
	"kaspar/repository"

	"github.com/gin-gonic/gin"
)

type StocksImplementation struct {
	Cache repository.Cache
}

func NewStocksImplementation(cache repository.Cache) *StocksImplementation {
	return &StocksImplementation{Cache: cache}
}

func (s *StocksImplementation) GetStocks(c *gin.Context) {

}

func (s *StocksImplementation) GetStocksToBuy(c *gin.Context) {

}

func (s *StocksImplementation) GetStocksToSell(c *gin.Context) {

}
