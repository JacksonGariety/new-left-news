package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Post struct {
	gorm.Model
	Title  string
	Url    string
	UserID uint
	User   User `gorm:"ForeignKey:UserID"`
}

type Posts []Post
