package models

import (
	"fmt"
	"sqlquerybuilder"

	_ "github.com/go-sql-driver/mysql"
)

var builder *sqlquerybuilder.SQLQueryBuilder

func ConnectDB() {
	fmt.Println("Connecting to database")

	builder = sqlquerybuilder.NewSQLQueryBuilder()
	builder.Connect()
}
func GetDBConnection() *sqlquerybuilder.SQLQueryBuilder {
	return builder
}
