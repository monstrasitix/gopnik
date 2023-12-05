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

func must(paths ...string) *template.Template {
	return template.Must(template.ParseFiles(paths...))
}

type ViewHandler struct {
    Name string
    Title string
}

func (v ViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

	Template[v.Name].ExecuteTemplate(w, "base", map[string]string{
        "lang": "en",
        "title": v.Title,
    })
}
