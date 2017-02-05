package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/go-zoo/bone"

	"github.com/JacksonGariety/new-left-news/app/models"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts := models.Posts{}
	utils.DB.Raw("SELECT x.* FROM posts x JOIN (SELECT p.id, SUM(v.vote) AS points FROM posts p JOIN points v ON v.post_id = p.id GROUP BY p.id) y ON y.id = x.id WHERE x.parent_post_id = 0 ORDER BY (y.points - 1)/POW(((EXTRACT(EPOCH FROM NOW()) - EXTRACT(EPOCH FROM x.created_at))/3600)+2, 1.5) DESC").Scan(&posts)
	posts.FetchUsers()
	posts.FetchPoints()
	utils.Render(w, r, "index.html", &utils.Props{
		"posts": posts,
	})
}

func Newest(w http.ResponseWriter, r *http.Request) {
	posts := models.Posts{}
	utils.DB.Select("id, title, url, user_id, created_at").Where("content = ?", "").Order("created_at DESC").Limit(30).Find(&posts)
	posts.FetchUsers()
	posts.FetchPoints()
	utils.Render(w, r, "index.html", &utils.Props{
		"posts": posts,
	})
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	if utils.DB.First(&post, id).RecordNotFound() {
		utils.NotFound(w, r)
	} else {
		utils.DB.Model(&post).Related(&post.User)
		utils.DB.Model(&post).Related(&post.Points)
		post.Posts = post.GetComments()
		utils.Render(w, r, "show_post.html", &utils.Props{
			"post": &post,
		})
	}
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "submit.html", &utils.Props{})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	form := utils.Props{
		"errors":   make(map[string]string),
		"title": r.FormValue("title"),
		"url": r.FormValue("url"),
	}

	if validateSubmitForm(form) == false {
		utils.Render(w, r, "submit.html", &form)
	} else {
		current_user := (*r.Context().Value("data").(*utils.Props))["current_user"]
		post := models.Post{
			Title: form["title"].(string),
			Url: form["url"].(string),
			UserID: current_user.(models.User).ID,
		}
		utils.DB.NewRecord(&post)
		utils.DB.Create(&post)
		post.UpvoteWithUser(current_user.(models.User))
		http.Redirect(w, r, "/newest", 307);
	}
}

func DestroyPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	utils.DB.First(&post, id)
	utils.DB.Model(&post).Related(&post.User)

	returnURL := "/"
	if post.ParentPostID != 0 {
		returnURL = fmt.Sprintf("/post/%d", post.ParentPostID)
	}

	current_user := (*r.Context().Value("data").(*utils.Props))["current_user"]
	post.DeleteWithUser(current_user.(models.User))
	http.Redirect(w, r, returnURL, 307);
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	form := utils.Props{
		"errors":   make(map[string]string),
		"content": r.FormValue("content"),
	}

	post := models.Post{}
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	utils.DB.First(&post, id)
	utils.DB.Model(&post).Related(&post.User)
	utils.DB.Model(&post).Related(&post.Points)
	post.Posts = post.GetComments()

	if validateCommentForm(form) == false {
		utils.Render(w, r, "show_post.html", &utils.Props{
			"post": &post,
			"errors": form["errors"],
			"content": form["content"],
		})
	} else {
		id, _ := strconv.Atoi(bone.GetValue(r, "id"))
		current_user := (*r.Context().Value("data").(*utils.Props))["current_user"]
		comment := models.Post{
			Content: form["content"].(string),
			UserID: current_user.(models.User).ID,
			ParentPostID: uint(id),
		}
		utils.DB.NewRecord(&comment)
		utils.DB.Create(&comment)
		http.Redirect(w, r, fmt.Sprintf("/post/%d", post.ID), 307)
	}
}

// Validations

func validateSubmitForm(form utils.Props) bool {
	form.ValidatePresence("title")
	if form.ValidatePresence("url") {
		form.ValidateUrl("url")
	}
	return form.IsValid()
}

func validateCommentForm(form utils.Props) bool {
	form.ValidatePresence("content")
	return form.IsValid()
}
