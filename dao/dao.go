package main

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func MysqlInit() *sql.DB {
	db, err := sql.Open("mysql", "root:syouZX@tcp(127.0.0.1:3306)/messageBoard?charset=utf8")
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	DB = db
	return DB
}
