package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nanpangyou/pixel-backend-go/internal/controller"
)

func New() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", controller.PingHandler)

	return r
}
