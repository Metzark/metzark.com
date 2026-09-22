package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Metzark/metzark.com/app/internal/handlers"
)

func main() {
	templates := template.Must(
		template.ParseFiles(
			"internal/templates/layout.html",
			"internal/templates/components/nav.html",
			"internal/templates/components/project-card.html",
			"internal/templates/components/footer.html",
			"internal/templates/home.html",
			"internal/templates/url-extender.html",
		),
	)

	mux := http.NewServeMux()

	// Static files
	mux.Handle(
		"GET /static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("./web/static")),
		),
	)

	// Home page
	homeHandler := &handlers.HomeHandler{
		Templates: templates,
	}

	urlExtenderHandler := &handlers.UrlExtenderHandler{
    	Templates: templates,
	}

	mux.HandleFunc("GET /", homeHandler.Handle)
	mux.HandleFunc("GET /url_extender", urlExtenderHandler.Handle)
	mux.HandleFunc("POST /url_extender", urlExtenderHandler.Handle)

	log.Println("server listening on http://localhost:3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}