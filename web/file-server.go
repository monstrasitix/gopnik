package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

func StaticAssets(router chi.Router) {
	route := "/static/*"
	strip := "/static"

	directory := http.Dir("./static")
	fileServer := http.FileServer(directory)

	router.Handle(
		route,
		http.StripPrefix(strip, fileServer))
}
