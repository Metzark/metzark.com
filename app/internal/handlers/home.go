package handlers

import (
	"html/template"
	"net/http"

	"github.com/Metzark/metzark.com/app/internal/models"
)

type HomeHandler struct {
	Templates *template.Template
}

type HomeData struct {
	Title    string
	Projects []models.Project
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	projects := []models.Project{
		{
			Name:        "Project One",
			Description: "Something cool I built.",
			URL:         "#",
			Icon:        "M",
			Status:      "Active",
			Color:       "#6c5ce7",
			Tags:        []string{"Go", "Nhost"},
		},
	}

	data := HomeData{
		Title:    "Metzark",
		Projects: projects,
	}

	err := h.Templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}