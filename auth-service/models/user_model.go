package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
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
	q := db_conn.GetTable("users").NewSelect()
	q.Where(fmt.Sprintf("id = %d", id))

	user := &User{}
	q.FindOne(&user)

	return user
}
func GetUserByUsername(name string) (*User, error) {
	db_conn := GetDBConnection()
	q := db_conn.GetTable("users").NewSelect()
	q.Where(fmt.Sprintf("username = \"%s\"", name))

	user := &User{}
	err := q.FindOne(
		&user.ID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// HashPassword generates a bcrypt hash for the given password.
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(bytes)
	return err
}

// VerifyPassword verifies if the given password matches the stored hash.
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
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
