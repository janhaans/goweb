package views

import "html/template"

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, filenames ...string) *View {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		panic(err)
	}
	t, err = t.ParseGlob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
