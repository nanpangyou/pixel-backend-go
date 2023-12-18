package cmd

import (
	"log"

	"github.com/nanpangyou/pixel-backend-go/internal/database"
	"github.com/nanpangyou/pixel-backend-go/internal/router"
	"github.com/spf13/cobra"
)

func Run() {

	rootCmd := &cobra.Command{
		Use: "pixel-backend-go",
	}

	dbCmd := &cobra.Command{
		Use: "db",
		// Run: func(cmd *cobra.Command, args []string) {
		// database.MySQLConnect()
		// database.MySQLCreatTable()
		// defer database.MySQLClose()
		// },
	}

	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.PgCreateTable()
		},
	}

	srvCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}

	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(srvCmd)
	dbCmd.AddCommand(createCmd)

	database.PgConnect()
	defer database.PgClose()
	rootCmd.Execute()

}

func RunServer() {
	r := router.New()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Panicln("Run 的下一行")
}
