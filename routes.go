package main

import (
	"net/http"
	"github.com/go-zoo/bone"

	"github.com/justinas/alice"
	"github.com/NYTimes/gziphandler"

	c "github.com/JacksonGariety/new-left-news/app/controllers"
	m "github.com/JacksonGariety/new-left-news/app/middleware"
)

func NewRouter() http.Handler {
	mux := bone.New()

	// Middleware
	chain := alice.New(
		m.Timeout,
		gziphandler.GzipHandler,
		m.Authenticate,
	)

	mux.Get("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Get("/", chain.ThenFunc(c.Index))
	mux.Get("/newest", chain.ThenFunc(c.Newest))
	mux.Get("/login", chain.Append(m.Retain).ThenFunc(c.LoginShow))
	mux.Post("/login", chain.Append(m.Retain).ThenFunc(c.Login))
	mux.Get("/logout", chain.ThenFunc(c.LogoutShow))
	mux.Get("/signup", chain.Append(m.Retain).ThenFunc(c.SignupShow))
	mux.Post("/signup", chain.Append(m.Retain).ThenFunc(c.SignupPost))
	mux.Get("/user/:name", chain.ThenFunc(c.UserShow))
	mux.Post("/user/:name", chain.ThenFunc(c.UserShow))

	return mux
}
