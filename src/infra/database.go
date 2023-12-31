package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	connStr := "postgresql://pixelpay:oot123@172.24.157.195:5432/pixelpay?sslmode=disable"

	if db != nil {
		return db, nil
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Connected to database!")
	return db, nil

}
