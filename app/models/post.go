package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

type Post struct {
	gorm.Model
	Title        string
	Url          string
	Content      string
	UserID       uint
	ParentPostID uint
	Points       Points
	Posts        Posts `gorm:"ForeignKey":ParentPostID`
	User         User `gorm:"ForeignKey:UserID"`
}

type Posts []Post

func (post *Post) CanDelete(current_user User) bool {
	return current_user.Name == post.User.Name
}

func (post *Post) CanUpvote(current_user User) bool {
	point := Point{
		UserID: current_user.ID,
		PostID: post.ID,
	}
	return current_user.Name == post.User.Name &&
		utils.DB.Where(&point).First(&point).RecordNotFound()
}

func (post *Post) DeleteWithUser(current_user User) {
	if post.CanDelete(current_user) {
		utils.DB.Unscoped().Delete(&post)
	}
}

func (post *Post) UpvoteWithUser(current_user User) {
	if post.CanUpvote(current_user) {
		point := Point{
			UserID: current_user.ID,
			PostID: post.ID,
		}
		utils.DB.NewRecord(&point)
		utils.DB.Create(&point)
	}
}

func (post *Post) FetchComments() {
	utils.DB.Select("*").Where("parent_post_id = ?", post.ID).Find(&post.Posts)
	post.Posts.FetchUsers()
}

func (posts Posts) FetchPoints() {
	for i, post := range posts {
		utils.DB.Model(&post).Related(&posts[i].Points)
	}
}

func (posts Posts) FetchUsers() {
	for i, post := range posts {
		utils.DB.Model(&post).Related(&posts[i].User)
	}
}
