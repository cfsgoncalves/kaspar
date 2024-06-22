package usecase

import (
	"kaspar/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemMonitoring struct {
	Cache repository.Cache
}

func NewSystemMonitoring(cache repository.Cache) *SystemMonitoring {
	return &SystemMonitoring{Cache: cache}
}

func (s *SystemMonitoring) Ping(context *gin.Context) {
	context.Status(http.StatusOK)
}

func (s *SystemMonitoring) Health(context *gin.Context) {
}
