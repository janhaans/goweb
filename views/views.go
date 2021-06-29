package views

import "html/template"

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, filenames ...string) *View {
	filenames = append(filenames, "views/layouts/footer.gohtml", "views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
