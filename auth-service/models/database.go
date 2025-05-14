package models

import "database/sql"

var db_conn *sql.DB

func ConnectDB() {

}

func GetDBConnection() *sql.DB {
	return db_conn
}
