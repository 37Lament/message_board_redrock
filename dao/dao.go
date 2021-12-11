package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB
//连接服务器

func InitDB() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test-b-g?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}

	dB = db
}
