package handler

import (
	"encoding/json"
	db "github.com/borntodie-new/information/pkg/db/sqlc"
	"github.com/justinas/nosurf"
	"log"
	"net/http"

	"github.com/borntodie-new/information/pkg/config"
	"github.com/borntodie-new/information/pkg/model"
	"github.com/borntodie-new/information/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

var Store db.Store

// Repository is the repository type
type Repository struct {
	App   *config.AppConfig
	Store db.Store
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository from the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// NewStore sets the store from the handlers
func NewStore(s db.Store) {
	Store = s
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	// sets current request's ip to session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "index.page.tmpl", &model.TemplateData{})
}

func (m *Repository) BackendLogin(w http.ResponseWriter, r *http.Request) {
	csrfToken := nosurf.Token(r)
	render.RenderTemplate(w, "backend.login.page.tmpl", &model.TemplateData{
		CSRFToken: csrfToken,
	})
}
func (m *Repository) BackendLoginForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := map[string]string{
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	}
	log.Println("提价表单：", r.Method)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (m *Repository) BackendIndex(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "backend.index.page.tmpl", &model.TemplateData{})
}

func (m *Repository) BackendUserCount(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.user.count.page.tmpl", &model.TemplateData{})
}

func (m *Repository) BackendUserList(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.user.list.page.tmpl", &model.TemplateData{})
}

func (m *Repository) BackendNewsEdit(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.news.edit.page.tmpl", &model.TemplateData{})
}
func (m *Repository) BackendNewsEditDetail(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.news.edit.detail.page.tmpl", &model.TemplateData{})
}
func (m *Repository) BackendNewsReview(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.news.review.page.tmpl", &model.TemplateData{})
}
func (m *Repository) BackendNewsReviewDetail(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.news.review.detail.page.tmpl", &model.TemplateData{})
}
func (m *Repository) BackendNewsType(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.news.type.page.tmpl", &model.TemplateData{})
}
