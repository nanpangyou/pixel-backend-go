package cmd

import (
	"log"

	"github.com/nanpangyou/pixel-backend-go/internal/router"
)

func RunServer() {
	log.Println("Server started================")
	r := router.New()
	log.Println("============Server started")
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Panicln("Run 的下一行")
}
