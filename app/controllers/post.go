package controllers

import (
	"net/http"

	"github.com/JacksonGariety/new-left-news/app/models"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts := new(models.Posts)
	utils.DB.Select("title, url, user_id").Find(posts)

	utils.Render(w, r, "index.html", &utils.Props{
		"posts": posts,
	})
}

func Newest(w http.ResponseWriter, r *http.Request) {
	posts := new(models.Posts)
	utils.DB.Select("title, url, user_id").Find(posts)

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
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	}

	if validateSubmitForm(form) == false {
		NewPost(w, r)
	} else {

	}
}

// Validations

func validateSubmitForm(form utils.Props) bool {
	return false
}
