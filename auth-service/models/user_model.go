package models

import (
	"database/sql"
	"fmt"
)

type UserModel struct {
	db_conn *sql.DB
}

type User struct {
	Username string
}

func GetUserModel() *UserModel {
	db_conn := GetDBConnection()
	m := &UserModel{
		db_conn: db_conn,
	}
	return m
}

func (m *UserModel) AddUser(user *User) {
	fmt.Println("Adding new user to database")
}
