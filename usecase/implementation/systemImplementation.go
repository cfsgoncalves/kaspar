package usecase

import (
	"kaspar/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemImplementation struct {
	Cache repository.Cache
}

func NewSystemImplementation(cache repository.Cache) *SystemImplementation {
	return &SystemImplementation{Cache: cache}
}

func (s *SystemImplementation) Ping(context *gin.Context) {
	context.Status(http.StatusOK)
}

func (s *SystemImplementation) Health(context *gin.Context) {
}
