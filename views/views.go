package views

import "html/template"

type View struct {
	Template *template.Template
}

func NewView(filenames ...string) *View {
	filenames = append(filenames, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		panic(err)
	}
	return &View{t}
}
