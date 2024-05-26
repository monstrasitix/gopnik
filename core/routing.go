package core

import (
	"github.com/go-chi/chi"
	apiV1 "github.com/monstrasitix/gopnik/api/v1"
	"github.com/monstrasitix/gopnik/web"
)

func Routing() *chi.Mux {
	router := chi.NewRouter()

	router.Group(web.Router)
    router.Route("/api/v1", apiV1.Router)

	return router
}
