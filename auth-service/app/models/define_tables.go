package models

import "sqlquerybuilder"

func DefineTables(builder *sqlquerybuilder.SQLQueryBuilder) {
	//define user table
	userTable := builder.NewTable("users")
	userTable.DefineColumn("id", "INT AUTO_INCREMENT PRIMARY KEY")
	userTable.DefineColumn("username", "VARCHAR(100) NOT NULL")
	userTable.DefineColumn("password", "VARCHAR(100) NOT NULL")
	userTable.EnsureTableExistsInDB()
	builder.DefineTable(userTable)

	//define api_key table
	sessionTable := builder.NewTable("sessions")
	sessionTable.DefineColumn("id", "INT AUTO_INCREMENT PRIMARY KEY")
	sessionTable.DefineColumn("user_id", "INT NOT NULL")
	sessionTable.DefineColumn("experation", "DATETIME NOT NULL")
	sessionTable.EnsureTableExistsInDB()
	builder.DefineTable(sessionTable)

	//define api_key table
	apiKeyTable := builder.NewTable("api_keys")
	apiKeyTable.DefineColumn("id", "INT AUTO_INCREMENT PRIMARY KEY")
	apiKeyTable.DefineColumn("api_key", "VARCHAR(200) NOT NULL")
	apiKeyTable.DefineColumn("owner_id", "INT NOT NULL")
	apiKeyTable.EnsureTableExistsInDB()
	builder.DefineTable(apiKeyTable)
}
