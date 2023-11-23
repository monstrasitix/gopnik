package web

import (
	"net/http"

	"github.com/go-chi/chi"
)


func Router(router chi.Router) {
	StaticAssets(router)
	ConfigureTemplates()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		View(w, "index", map[string]string{
			"lang": "en",
			"title": "Homepage",
		})
	})
}
