package web

import (
	"html/template"
	"net/http"
)

var (
	Template = make(map[string]*template.Template)
)

const (
	LAYOUT_MAIN = "./web/view/layout/main.html"
)

func ConfigureTemplates() {
	Template["index"] = must(LAYOUT_MAIN, "./web/view/index.html")
}

func View(w http.ResponseWriter, name string, payload any) {
	w.Header().Set("Content-Type", "text/html")

	Template[name].ExecuteTemplate(w, "base", payload)
}

func must(paths ...string) *template.Template {
	return template.Must(template.ParseFiles(paths...))
}
