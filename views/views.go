package views

import (
	"html/template"
	"net/http"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func NewView(layout string, filenames ...string) *View {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		panic(err)
	}
	t, err = t.ParseGlob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
