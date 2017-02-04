package controllers

import (
	"net/http"
	"strconv"
	"github.com/go-zoo/bone"

	"github.com/JacksonGariety/new-left-news/app/models"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts := models.Posts{}
	utils.DB.Select("id, title, url, user_id").Limit(30).Find(&posts)
	posts.FetchUsers()
	utils.Render(w, r, "index.html", &utils.Props{
		"posts": posts,
	})
}

func Newest(w http.ResponseWriter, r *http.Request) {
	posts := models.Posts{}
	utils.DB.Select("id, title, url, user_id").Limit(30).Find(&posts)
	posts.FetchUsers()
	utils.Render(w, r, "index.html", &utils.Props{
		"posts": posts,
	})
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
		http.Redirect(w, r, "/newest", 307);
	}
}

func DestroyPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	utils.DB.First(&post, id)
	utils.DB.Model(&post).Related(&post.User)
	authorized_username := (*r.Context().Value("data").(*utils.Props))["authorized_username"]
	post.DeleteWithUser(authorized_username.(string))
	http.Redirect(w, r, "/", 307);
}

// Validations

func validateSubmitForm(form utils.Props) bool {
	form.ValidatePresence("title")
	if form.ValidatePresence("url") {
		form.ValidateUrl("url")
	}
	return form.IsValid()
}
