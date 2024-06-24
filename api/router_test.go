package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

type StocksMock struct {
}

func (s *StocksMock) GetStockByName(date string, stockName string) (any, error) {
	return nil, nil
}

func TestGetStockByNameAndOptionalDate(t *testing.T) {
	os.Setenv("REDIS_PASSWORD", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	os.Setenv("DB", "0")
	os.Setenv("REDIS_TTL", "10")
	os.Setenv("REDIS_SERVER", "localhost")
	os.Setenv("REDIS_PORT", "6379")

	//Add output validation
	t.Run("happy_path_integration", func(t *testing.T) {
		router := NewRouter()

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/v1/stocks/TSLA/2024-06-24", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		//assert.Equal(t, "", w.Body.String())

	})

	t.Run("test_with_date", func(t *testing.T) {
		router := NewRouter()

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/v1/stocks/TSLA/2024-06-24", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

	})

	t.Run("test_without_date", func(t *testing.T) {
		router := NewRouter()

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/v1/stocks/TSLA/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

func TestHealth(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {

	})
}

func TestPing(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {

	})
}
