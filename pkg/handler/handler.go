package handler

import (
	"net/http"

	"github.com/borntodie-new/information/pkg/config"
	"github.com/borntodie-new/information/pkg/model"
	"github.com/borntodie-new/information/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository from the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	// sets current request's ip to session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "index.page.tmpl", &model.TemplateData{})
}

func (m *Repository) BackendLogin(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "backend.login.page.tmpl", &model.TemplateData{})
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
