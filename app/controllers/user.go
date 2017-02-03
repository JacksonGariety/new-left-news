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

func NewUser(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "signup.html", &utils.Props{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	form := utils.Props{
		"errors":                make(map[string]string),
		"email":                 r.FormValue("email"),
		"username":              r.FormValue("username"),
		"password":              r.FormValue("password"),
		"password_confirmation": r.FormValue("password_confirmation"),
	}

	if validateSignupForm(form) == false {
		utils.Render(w, r, "signup.html", &form)
	} else {
		(&models.User{Name: form["username"].(string), Email: form["email"].(string)}).CreateFromPassword(form["password"].(string))
		signedToken, expireCookie, claims := models.ClaimsCreate(form["username"].(string)) // creates a JWT token
		cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, claims.Userpath(), 307)
	}
}

// Validations

func validateSignupForm(form utils.Props) bool {
	if form.ValidatePresence("email") {
		if form.ValidateEmail("email") {
			user := &models.User{Email: form["email"].(string)}
			exists := !utils.DB.Where(user).First(user).RecordNotFound()
			if exists {
				form.SetError("email", "email is already in use")
			}
		}
	}

	if form.ValidatePresence("password") {
		form.ValidateLength("password", 5, 30)
	}

	if form.ValidatePresence("username") {
		form.ValidateNoSpace("username")

		user := &models.User{Name: form["username"].(string)}
		exists := !utils.DB.Where(user).First(user).RecordNotFound()
		if exists {
			form.SetError("username", "username is already in use")
		}
	}

	if form.ValidatePresence("password_confirmation") {
		form.ValidateConfirmation("password", "password_confirmation")
	}

	return form.IsValid()
}
