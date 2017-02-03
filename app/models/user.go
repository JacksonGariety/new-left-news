package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

type User struct {
	gorm.Model
	Email        string
	Name         string
	PasswordHash string
	Admin        bool
}

type Users []User

func (user *User) CreateFromPassword(password string) (*User, error) {
	user.PasswordHash = hashPassword(password)
	utils.DB.NewRecord(user)
	c := utils.DB.Create(&user)
	return user, c.Error
}

func (user *User) Userpath() string {
	return fmt.Sprintf("/user/%s", user.Name)
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
