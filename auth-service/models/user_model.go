package models

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Password string
}

func NewUser() *User {
	m := &User{}
	return m
}
func GetUserByID(id int) *User {
	db_conn := GetDBConnection()
	q := db_conn.GetTable("users").NewQuery()
	q.Where(fmt.Sprintf("id = %d", id))

	user := &User{}
	q.FindOne(&user)

	return user
}
func GetUserByUsername(name string) *User {
	db_conn := GetDBConnection()
	q := db_conn.GetTable("users").NewQuery()
	q.Where(fmt.Sprintf("username = %s", name))

	user := &User{}
	q.FindOne(&user)

	return user
}

func (user *User) Save() {
	fmt.Println("Adding new user to database")
	db_conn := GetDBConnection()
	q := db_conn.GetTable("users").NewInsert()

	//Add data to insert statement
	q.AddIntColumn("id", user.ID)
	q.AddStringColumn("username", user.Username)
	q.AddStringColumn("password", user.Password)

	q.Send()
}
