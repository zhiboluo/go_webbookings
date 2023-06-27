package handlers

import (
	"net/http"

	"github.com/zhiboluo/go_webbookings/pkg/config"
	"github.com/zhiboluo/go_webbookings/pkg/models"
	"github.com/zhiboluo/go_webbookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// handlers is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello nihao."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	// send the data to the template

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
