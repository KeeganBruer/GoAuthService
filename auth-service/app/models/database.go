package models

import (
	"fmt"
	"sqlquerybuilder"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Models struct {
	builder *sqlquerybuilder.SQLQueryBuilder
}
type BaseModel struct {
	models *Models
}

func (model *BaseModel) GetDBQueryBuilder() *sqlquerybuilder.SQLQueryBuilder {
	return model.models.builder
}

type DBConnection struct {
	User   string
	Passwd string
	Addr   string
	Name   string
}

func ConnectDB(connConfig DBConnection) *Models {
	fmt.Println("Connecting to database")
	cfg := mysql.NewConfig()
	cfg.User = connConfig.User
	cfg.Passwd = connConfig.Passwd
	cfg.Net = "tcp"
	cfg.Addr = connConfig.Addr

	builder := sqlquerybuilder.NewSQLQueryBuilder()
	builder.Connect(cfg)
	builder.UseDatabase(connConfig.Name)
	DefineTables(builder)
	models := &Models{
		builder: builder,
	}
	return models
}
