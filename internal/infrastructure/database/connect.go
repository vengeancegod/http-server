package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DB")

	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
	state, err := sql.Open("mysql", connect)
	if err != nil {
		log.Fatalf("Failed to connect DB", err)
		return nil, err
	}

	return state, nil
}
