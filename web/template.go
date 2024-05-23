package web

import (
	"html/template"
)

var (
	Template = make(map[string]*template.Template)
)

const (
	LAYOUT_MAIN = "./web/view/layout/main.html"
)

func must(paths ...string) *template.Template {
    return template.Must(template.ParseFiles(paths...))
}

func ConfigureTemplates() {
	Template["index"] = must(LAYOUT_MAIN, "./web/view/index.html")
	Template["about"] = must(LAYOUT_MAIN, "./web/view/about.html")
}


