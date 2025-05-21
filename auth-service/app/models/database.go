package models

import (
	"fmt"
	"os"
	"sqlquerybuilder"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Models struct {
	builder *sqlquerybuilder.SQLQueryBuilder
}

func ConnectDB() *Models {
	fmt.Println("Connecting to database")
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUser")
	cfg.Passwd = os.Getenv("DBPass")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBAdrs")
	//cfg.DBName = os.Getenv("DBName")
	builder := sqlquerybuilder.NewSQLQueryBuilder()
	builder.Connect(cfg)
	builder.UseDatabase(os.Getenv("DBName"))
	DefineTables(builder)
	models := &Models{
		builder: builder,
	}
	return models
}
