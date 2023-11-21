package cmd

import (
	"log"

	"github.com/nanpangyou/pixel-backend-go/internal/database"
	"github.com/nanpangyou/pixel-backend-go/internal/router"
)

func RunServer() {
	database.PgConnect()
	database.PgCreateTable()
	defer database.PgClose()
	database.MySQLConnect()
	database.MySQLCreatTable()
	defer database.MySQLClose()
	r := router.New()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Panicln("Run 的下一行")
}
