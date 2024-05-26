package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Page struct {
	Language string
	Title    string
}

func view(name string, data map[string]any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Template[name].ExecuteTemplate(w, "base", data)
	}
}

func Router(router chi.Router) {
	StaticAssets(router)
	ConfigureTemplates()

	router.Get("/", view("index", map[string]any{
		"lang":  "en",
		"title": "Homepage",
	}))

	router.Get("/about", view("about", map[string]any{
		"lang":  "en",
		"title": "About",
	}))
}
