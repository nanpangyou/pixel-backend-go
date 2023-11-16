package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong1")
	c.Writer.WriteString("pong2")
}
