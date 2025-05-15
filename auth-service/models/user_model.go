package models

import (
	"fmt"
)

type User struct {
	Username string
}

func NewUser() *User {
	m := &User{}
	return m
}
func GetUserByID(id int) *User {
	db_conn := GetDBConnection()
	q := db_conn.GetTable("users").NewQuery()
	q.Where("")

	user := &User{}
	q.FindOne(&user)

	return user
}

func (m *User) Save() {
	fmt.Println("Adding new user to database")
	db_conn := GetDBConnection()
	q := db_conn.GetTable("users").NewInsert()
	q.Send()
}
