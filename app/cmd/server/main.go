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
			"internal/templates/home.html",
			"internal/templates/components/nav.html",
			"internal/templates/components/project-card.html",
			"internal/templates/components/footer.html",
		),
	)

	mux := http.NewServeMux()

	// Static files
	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("./web/static")),
		),
	)

	// Home page
	homeHandler := &handlers.HomeHandler{
		Templates: templates,
	}

	mux.Handle("/", homeHandler)

	log.Println("server listening on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}