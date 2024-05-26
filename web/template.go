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

func ConfigureTemplates() {

	Template["index"] = template.Must(template.ParseFiles(
		LAYOUT_MAIN,
		"./web/view/index.html"))

	Template["index"] = template.Must(template.ParseFiles(
		LAYOUT_MAIN,
		"./web/view/about.html"))
}
