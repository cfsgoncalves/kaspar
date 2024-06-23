package api

import (
	"kaspar/configuration"
	usecase "kaspar/usecase/interface"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type StockAPI struct {
	StockAPI usecase.Stocks
}

func NewStockApi(stockApi usecase.Stocks) *StockAPI {
	return &StockAPI{StockAPI: stockApi}
}

func (s *StockAPI) GetStockByNameAndOptionalDate(c *gin.Context) {
	dateParam, _ := c.Params.Get("date")
	dateParam = strings.Trim(dateParam, "/")
	stockName, _ := c.Params.Get("name")
	DATE_FORMAT := configuration.GetEnvAsString("DATE_FORMAT", "2006-01-02")
	date := time.Now().UTC().Format(DATE_FORMAT)

	//Validate date paramenter
	if dateParam != "" {
		_, err := time.Parse(DATE_FORMAT, dateParam)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		date = dateParam
		println(date)
	}

	json, err := s.StockAPI.GetStockByName(date, stockName)

	switch err {
	case nil:
		c.JSON(http.StatusOK, json)
	default:
		c.Status(http.StatusInternalServerError)
	}
}
