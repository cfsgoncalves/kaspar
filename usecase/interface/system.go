package usecase

import "github.com/gin-gonic/gin"

type System interface {
	Ping(*gin.Context)
	Health(*gin.Context)
}
