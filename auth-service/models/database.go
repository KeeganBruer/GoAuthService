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
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	builder = sqlquerybuilder.NewSQLQueryBuilder()
	builder.Connect(cfg)

	//define user table
	userTable := builder.NewTable("users")
	userTable.DefineColumn("id", "INT AUTO_INCREMENT PRIMARY KEY")
	userTable.DefineColumn("username", "VARCHAR(100) NOT NULL")
	userTable.DefineColumn("password", "VARCHAR(100) NOT NULL")
	userTable.EnsureTableExistsInDB()
	builder.DefineTable(userTable)

	//define api_key table
	apiKeyTable := builder.NewTable("api_key")
	apiKeyTable.DefineColumn("id", "INT AUTO_INCREMENT PRIMARY KEY")
	apiKeyTable.DefineColumn("key", "VARCHAR(200) NOT NULL")
	apiKeyTable.EnsureTableExistsInDB()
	builder.DefineTable(apiKeyTable)

}
func GetDBConnection() *sqlquerybuilder.SQLQueryBuilder {
	return builder
}
