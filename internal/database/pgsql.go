package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // 如果不引入驱动，则无法连接数据库
)

const (
	pgHost     = "pg-for-pixel-backend"
	pgPort     = 5432
	pgUser     = "pixel"
	pgPassword = "123456"
	pgDbname   = "pixel_dev"
)

func PgConnect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pgHost, pgPort, pgUser, pgPassword, pgDbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	PgDB = db
	// 检查数据库连接是否成功
	if err = PgDB.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("数据库连接成功")

	return PgDB, nil
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
	_, err := PgDB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("创建表成功")

}

func PgMigrate() {
	// 修改表的 SQL 语句
	migrateTableSQL := `
		ALTER TABLE users
		ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
	`
	// 执行创建表的 SQL 语句
	_, err := PgDB.Exec(migrateTableSQL)

	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("修改表成功")
}

func PgClose() {
	PgDB.Close()
	log.Println("数据库连接关闭")
}
