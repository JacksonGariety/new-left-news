package main

import (
	"net/http"
	"github.com/go-zoo/bone"

	"github.com/justinas/alice"
	"github.com/NYTimes/gziphandler"

	c "github.com/JacksonGariety/new-left-news/app/controllers"
	m "github.com/JacksonGariety/new-left-news/app/middleware"
	"github.com/JacksonGariety/new-left-news/app/utils"
)

func NewRouter() http.Handler {
	mux := bone.New()

	// Middleware
	chain := alice.New(
		m.Timeout,
		gziphandler.GzipHandler,
		m.Authenticate,
	)

	mux.NotFoundFunc(utils.NotFound)
	mux.Get("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Get("/favicon.ico", http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "static/favicon.ico") }))
	mux.Get("/", chain.ThenFunc(c.Index))
	mux.Get("/newest", chain.ThenFunc(c.Newest))
	mux.Post("/newest", chain.ThenFunc(c.Newest))
	mux.Get("/login", chain.Append(m.Retain).ThenFunc(c.NewSession))
	mux.Post("/login", chain.Append(m.Retain).ThenFunc(c.CreateSession))
	mux.Get("/logout", chain.Append(m.Protect).ThenFunc(c.DestroySession))
	mux.Get("/signup", chain.Append(m.Retain).ThenFunc(c.NewUser))
	mux.Post("/signup", chain.Append(m.Retain).ThenFunc(c.CreateUser))
	mux.Get("/user/:name", chain.ThenFunc(c.ShowUser))
	mux.Post("/user/:name", chain.ThenFunc(c.ShowUser))
	mux.Get("/submit", chain.Append(m.Protect).ThenFunc(c.NewPost))
	mux.Post("/submit", chain.Append(m.Protect).ThenFunc(c.CreatePost))
	mux.Get("/post/:id", chain.ThenFunc(c.ShowPost))
	mux.Post("/post/:id", chain.ThenFunc(c.ShowPost))
	mux.Get("/post/:id/submit", chain.Append(m.Protect).ThenFunc(c.ShowPost))
	mux.Post("/post/:id/submit", chain.Append(m.Protect).ThenFunc(c.CreateComment))
	mux.Get("/post/:id/delete", chain.Append(m.Protect).ThenFunc(c.DestroyPost))
	mux.Get("/post/:id/upvote", chain.Append(m.Protect).ThenFunc(c.UpvotePost))

	return mux
}
