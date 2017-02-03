package controllers

import (
	"github.com/go-zoo/bone"
	"net/http"

	"github.com/JacksonGariety/new-left-news/app/models"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

func UserShow(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: bone.GetValue(r, "name")}
	if !utils.DB.Where(user).First(user).RecordNotFound() {
		utils.Render(w, r, "user.html", &utils.Props{
			"username": user.Name,
			"admin":    user.Admin,
		})
	} else {
		utils.NotFound(w, r)
	}
}
