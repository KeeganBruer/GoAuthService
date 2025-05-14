package models

import (
	"database/sql"
)

type UserModel struct {
	db_conn *sql.DB
}

func GetUserModel() *UserModel {
	db_conn := GetDBConnection()
	m := &UserModel{
		db_conn: db_conn,
	}
	return m
}

func (m *UserModel) AddUser() {

}
