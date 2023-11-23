package core

import (
	"log"
	"net/http"
)

func NewServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
}

func Run(server *http.Server) {
	log.Printf("Listening on: http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
