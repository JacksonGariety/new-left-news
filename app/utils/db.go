package utils

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB(dbstring string) {
	if dbstring == "$dbstringnln" {
	  dbstring = os.Getenv("dbstringnln")
	}

	var err error
	DB, err = gorm.Open("postgres", dbstring)

	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}
