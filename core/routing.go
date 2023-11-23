package core

import (
	"github.com/go-chi/chi"
	"github.com/monstrasitix/gopnik/web"
)

func Routing() *chi.Mux {
	router := chi.NewRouter()

	router.Group(web.Router)

	return router
}
