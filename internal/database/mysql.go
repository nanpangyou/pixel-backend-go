package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlHost     = "mysql-for-pixel-backend"
	mysqlPort     = 3306
	mysqlUser     = "pixel"
	mysqlPassword = "123456"
	mysqlDbname   = "pixel_dev"
)

func MySQLConnect() {
	mysqlConnectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDbname)

	if db, err := sql.Open("mysql", mysqlConnectString); err != nil {
		panic(err.Error())
	} else {
		MySQLDB = db
		log.Println("mysql连接成功")
	}
}

func MySQLCreatTable() {
	// 创建 "user" 表
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS user (
  			id INT AUTO_INCREMENT PRIMARY KEY,
 			name VARCHAR(255) NOT NULL,
		    email VARCHAR(255) NOT NULL UNIQUE,
		    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := MySQLDB.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}

	log.Println("mysql User table created successfully!")
}

func MySQLClose() {
	MySQLDB.Close()
	log.Println("mysql 关闭成功")
}
