package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysql Tutorial")

	db, err := sql.Open("mysql", "root:adminadmin@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO test VALUES (2, 'TEST')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
