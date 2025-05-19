# SQL Query Builder
Currently Supported Drivers: MySQL

## Usage
1. Define configuration and connect to database 
```golang
cfg := mysql.NewConfig()
cfg.User = os.Getenv("DBUser")
cfg.Passwd = os.Getenv("DBPass")
cfg.Net = "tcp"
cfg.Addr = os.Getenv("DBAdrs")
cfg.DBName = os.Getenv("DBName")

builder = sqlquerybuilder.NewSQLQueryBuilder()
builder.Connect(cfg)
```

2. Define a table 
```golang
//define user table
userTable := builder.NewTable("users")
userTable.DefineColumn("id", "INT AUTO_INCREMENT PRIMARY KEY")
userTable.DefineColumn("username", "VARCHAR(100) NOT NULL")
userTable.DefineColumn("password", "VARCHAR(100) NOT NULL")
userTable.EnsureTableExistsInDB()
builder.DefineTable(userTable) //Allows access from builder.GetTable()

```

3. Create a new insert statement for a given table
```golang
q := builder.GetTable("users").NewInsert()
//Add data to insert statement
q.AddIntColumn("id", user.ID)
q.AddStringColumn("username", user.Username)
q.AddStringColumn("password", user.Password)
q.Send()

```

4. Execute a select statement
```golang
q := builder.GetTable("users").NewSelect()
q.Where(fmt.Sprintf("username = %s", name))

//Find One
user := &User{}
q.FindOne(&user)

//Find All
var users []*User
q.FindAll(func(get func(dest ...any) error) error {
    user := &User{}
    err := get(
        &user.ID,
        &user.Username,
        &user.Password,
    )
    users = append(users, user)
    return err
})
```