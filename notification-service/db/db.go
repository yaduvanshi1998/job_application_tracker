package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() {
	// Define MySQL connection credentials
	user := "root"          // your MySQL username
	password := "Crime_123" // your MySQL password
	host := "localhost"     // MySQL host
	port := "3306"          // MySQL port
	dbname := "job_tracker" // database name

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	// Connect to MySQL using sqlx
	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("DB connection error:", err)
	}

}
