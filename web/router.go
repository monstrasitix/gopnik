package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/monstrasitix/gopnik/database/user"
)

type Page struct {
    Language string
    Title string
}


func Router(router chi.Router) {
	StaticAssets(router)
	ConfigureTemplates()

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        users := user.GetUsers()

        w.Header().Set("Content-Type", "text/html")

        Template["index"].ExecuteTemplate(w, "base", map[string]any{
            "lang": "en",
            "title": "Homepage",
            "users": users,
        })
    })
}
