package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:id`
	Name string `json:name`
}

func main() {
	fmt.Println("Go Mysql Tutorial")

	db, err := sql.Open("mysql", "root:adminadmin@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO test VALUES (1, 'TEST 1')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	results, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag

		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}

		log.Printf(tag.Name)
	}
}
