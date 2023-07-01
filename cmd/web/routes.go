package main

import (
	"github.com/borntodie-new/information/pkg/config"
	"github.com/borntodie-new/information/pkg/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.Abort))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)

	mux.Get("/backend/login", handler.Repo.BackendLogin)
	mux.Post("/backend/login_form", handler.Repo.BackendLoginForm)

	mux.Get("/backend/index", handler.Repo.BackendIndex)
	mux.Get("/backend/user_count", handler.Repo.BackendUserCount)
	mux.Get("/backend/user_list", handler.Repo.BackendUserList)
	mux.Get("/backend/news_edit", handler.Repo.BackendNewsEdit)
	mux.Get("/backend/news_edit_detail/{id}", handler.Repo.BackendNewsEditDetail)
	mux.Get("/backend/news_review", handler.Repo.BackendNewsReview)
	mux.Get("/backend/news_review_detail/{id}", handler.Repo.BackendNewsReviewDetail)
	mux.Get("/backend/news_type", handler.Repo.BackendNewsType)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
