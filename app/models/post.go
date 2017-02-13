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
	return current_user.Name != post.User.Name &&
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
			Vote: 1,
			UserID: current_user.ID,
			PostID: post.ID,
		}
		utils.DB.NewRecord(&point)
		utils.DB.Create(&point)
	}
}

// recursively gets all comments including nested
func (post Post) FetchAllComments() Posts {
	comments := Posts{}
	utils.DB.Select("*").Where("parent_post_id = ?", post.ID).Find(&comments)
	comments.FetchPoints()
	comments.FetchUsers()
	for i, comment := range comments {
		comments[i].Posts = comment.FetchAllComments()
	}
	return comments
}

func (posts Posts) FetchPoints() {
	for i, post := range posts {
		utils.DB.Model(&post).Related(&posts[i].Points)
	}
}

// only returns one level of comments
func (posts Posts) FetchComments() {
	for i, post := range posts {
		comments := Posts{}
		utils.DB.Select("*").Where("parent_post_id = ?", post.ID).Find(&comments)
		posts[i].Posts = comments
	}
}

func (posts Posts) FetchUsers() {
	for i, post := range posts {
		utils.DB.Model(&post).Related(&posts[i].User)
	}
}
