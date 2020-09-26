package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

//Connect connects to database
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to db")
	con = db
	return db
}
