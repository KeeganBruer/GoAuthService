package models

import (
	"fmt"
	"os"
	"sqlquerybuilder"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var builder *sqlquerybuilder.SQLQueryBuilder

func ConnectDB() {
	fmt.Println("Connecting to database")
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUser")
	cfg.Passwd = os.Getenv("DBPass")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBAdrs")
	//cfg.DBName = os.Getenv("DBName")

	builder = sqlquerybuilder.NewSQLQueryBuilder()
	builder.Connect(cfg)
	builder.UseDatabase(os.Getenv("DBName"))
	DefineTables(builder)
}
func GetDBConnection() *sqlquerybuilder.SQLQueryBuilder {
	return builder
}
