package controllers

import (
	"net/http"
	"strconv"
	"github.com/go-zoo/bone"

	"github.com/JacksonGariety/new-left-news/app/models"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

func UpvotePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	utils.DB.First(&post, id)
	utils.DB.Model(&post).Related(&post.User)
	current_user := (*r.Context().Value("data").(*utils.Props))["current_user"]
	post.UpvoteWithUser(current_user.(models.User))
	http.Redirect(w, r, "/", 307);
}
