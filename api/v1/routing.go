package apiV1

import (
	"github.com/go-chi/chi"
	"github.com/monstrasitix/gopnik/api/v1/handler"
)

func Router(router chi.Router) {
	router.Get("/products", handler.GetProducts)
	router.Get("/products/{id}", handler.GetProduct)
}
