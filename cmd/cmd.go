package cmd

import (
	"log"

	"github.com/nanpangyou/pixel-backend-go/internal/database"
	"github.com/nanpangyou/pixel-backend-go/internal/router"
)

func RunServer() {
	// defer database.PgClose(database.DB)
	database.PgConnect()
	database.PgCreateTable()
	r := router.New()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Panicln("Run 的下一行")
}
