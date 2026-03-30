package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:Crime_123@tcp(127.0.0.1:3306)/job_tracker?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected successfully:", db)
}
