package web

import (
	"net/http"

	"github.com/go-chi/chi"
)


func Router(router chi.Router) {
	StaticAssets(router)
	ConfigureTemplates()

	router.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		View(w, "about", nil)
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		View(w, "index", nil)
	})
}
