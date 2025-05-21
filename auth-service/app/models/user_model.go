package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	models *Models
}

func (models *Models) GetUserModel() *UserModel {
	return &UserModel{
		models: models,
	}
}

type User struct {
	models   *UserModel
	ID       int
	Username string
	Password string
}

func (userModel *UserModel) NewUser() *User {
	m := &User{
		models: userModel,
	}
	return m
}
func (userModel *UserModel) GetUserByID(id int) (*User, error) {
	builder := userModel.models.builder
	q := builder.GetTable("users").NewSelect()
	q.Where(fmt.Sprintf("id = %d", id))

	user := &User{
		models: userModel,
	}
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
func (userModel *UserModel) GetUserByUsername(name string) (*User, error) {
	builder := userModel.models.builder
	q := builder.GetTable("users").NewSelect()
	q.Where(fmt.Sprintf("username = \"%s\"", name))

	user := &User{
		models: userModel,
	}
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
	builder := user.models.models.builder
	q := builder.GetTable("users").NewInsert()

	//Add data to insert statement
	q.AddIntColumn("id", user.ID)
	q.AddStringColumn("username", user.Username)
	q.AddStringColumn("password", user.Password)

	q.Send()
}
