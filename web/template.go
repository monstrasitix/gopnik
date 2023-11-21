package web

import (
	"html/template"
	"io"
)

var (
	Template = make(map[string]*template.Template)
)

const (
	LAYOUT_MAIN = "./web/view/layout/main.html"
)

func ConfigureTemplates() {
	Template["index"] = must(LAYOUT_MAIN, "./web/view/index.html")
	Template["about"] = must(LAYOUT_MAIN, "./web/view/about.html")
}

func View(w io.Writer, name string, payload any) {
	Template[name].ExecuteTemplate(w, "base", payload)
}

func must(paths ...string) *template.Template {
	return template.Must(template.ParseFiles(paths...))
}