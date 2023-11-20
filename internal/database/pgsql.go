package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // 如果不引入驱动，则无法连接数据库
)

const (
	host     = "pg-for-pixel-backend"
	port     = 5432
	user     = "pixel"
	password = "123456"
	dbname   = "pixel_dev"
)

func PgConnect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	DB = db
	// 检查数据库连接是否成功
	if err = DB.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("数据库连接成功")

	return DB, nil
}

func PgCreateTable() {
	// 创建表的 SQL 语句
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(50) NOT NULL,
            email VARCHAR(100) NOT NULL,
            password VARCHAR(100) NOT NULL
        );
    `
	// 执行创建表的 SQL 语句
	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("创建表成功")

}

func PgClose() {
	DB.Close()
	log.Println("数据库连接关闭")
}
