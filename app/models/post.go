package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

type Post struct {
	gorm.Model
	Title  string
	Url    string
	UserID uint
	User   User `gorm:"ForeignKey:UserID"`
}

type Posts []Post

func (post *Post) CanDelete(authorized_username string) bool {
	return authorized_username == post.User.Name
}

func (post *Post) DeleteWithUser(authorized_username string) {
	if post.CanDelete(authorized_username) {
		utils.DB.Unscoped().Delete(&post)
	}
}

func (posts Posts) FetchUsers() {
	for i, post := range posts {
		utils.DB.Model(&post).Related(&posts[i].User)
	}
}
