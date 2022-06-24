package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:hcy990202@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	id       int
	username string
	password string
}

// 查询单行数据
func queryOneRow() {
	s := "select * from user_tbl where id = ?"
	var u User
	err := db.QueryRow(s, 2).Scan(&u.id, &u.username, &u.password) //在db.QueryRow()中设置查询的行数
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("u: %v\n", u)
	}
}

// 查询多行数据
func queryManyRow() {
	s := "select * from user_tbl"
	r, err := db.Query(s)
	var u User
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for r.Next() {
			r.Scan(&u.id, &u.username, &u.password)
			fmt.Printf("u: %v\n", u)
		}
	}
}

func main() {

	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("链接成功")
	}

	// queryOneRow()

	queryManyRow()
}
