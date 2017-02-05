package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Point struct {
	gorm.Model
	UserID uint
	PostID uint
	Vote   int
	User   User `gorm:"ForeignKey:UserID"`
	Post   Post `gorm:"ForeignKey:PostID"`
}

type Points []Point
