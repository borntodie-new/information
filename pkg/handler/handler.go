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

//func (m *Repository) Abort(w http.ResponseWriter, r *http.Request) {
//	// get the last request's ip in here
//	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
//
//	stringMap := make(map[string]string)
//	stringMap["test"] = "Hello, again"
//	stringMap["remote_ip"] = remoteIP
//	render.RenderTemplate(w, "about.page.tmpl", &model.TemplateData{
//		StringMap: stringMap,
//	})
//}
