# job_application_tracker
While applying to many jobs, I needed a way to track applications, statuses, and follow-ups in one place, so I built a Job Application Tracker using React and Golang


1. Initialize Go Module
cd backend
go mod init job-tracker-backend
go get github.com/gin-gonic/gin
go get github.com/jmoiron/sqlx       # optional for DB
go get github.com/mattn/go-sqlite3  # if using SQLite

2. test sql connection

package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:password@tcp(127.0.0.1:3306)/job_tracker?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected successfully:", db)