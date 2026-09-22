package handlers

import (
	"html/template"
	"net/http"
)

type UrlExtenderHandler struct {
	Templates *template.Template
}

type UrlExtenderData struct {
	Title string
}

func (h *UrlExtenderHandler) Handle(w http.ResponseWriter, r *http.Request) {
	data := UrlExtenderData{
		Title: "URL Extender",
	}

	if err := h.Templates.ExecuteTemplate(w, "url-extender", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}