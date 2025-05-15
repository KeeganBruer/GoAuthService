package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db_conn *sql.DB

func ConnectDB() {
	fmt.Println("Connecting to database")
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	db_conn = db
}
func CloseDB() {
	db_conn.Close()
}
func GetDBConnection() *sql.DB {
	return db_conn
}
