package cmd

import (
	"github.com/nanpangyou/pixel-backend-go/internal/router"
)

func RunServer() {
	r := router.New()
	r.Run(":8080")
}
