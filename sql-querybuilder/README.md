# SQL Query Builder

```golang
builder = sqlquerybuilder.NewSQLQueryBuilder()
builder.Connect()

//define user table
userTable := builder.NewTable("users")
userTable.DefineColumn("id", "INT AUTO_INCREMENT PRIMARY KEY")
userTable.DefineColumn("username", "VARCHAR(100) NOT NULL")
userTable.DefineColumn("password", "VARCHAR(100) NOT NULL")
userTable.EnsureTableExistsInDB()
builder.DefineTable(userTable)


q := builder.GetTable("users").NewInsert()
//Add data to insert statement
q.AddIntColumn("id", user.ID)
q.AddStringColumn("username", user.Username)
q.AddStringColumn("password", user.Password)
q.Send()


q := builder.GetTable("users").NewQuery()
q.Where(fmt.Sprintf("username = %s", name))
user := &User{}
q.FindOne(&user)

```