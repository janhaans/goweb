package controllers

import (
	"fmt"
	"net/http"

	"github.com/janhaans/goweb/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Email = %s\n", r.PostForm["email"])
	fmt.Fprintf(w, "Password = %s\n", r.PostForm["password"])
}
