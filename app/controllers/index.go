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
