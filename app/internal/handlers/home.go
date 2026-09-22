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

func (h *HomeHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

	if err := h.Templates.ExecuteTemplate(w, "home", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}