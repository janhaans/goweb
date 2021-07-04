package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/janhaans/goweb/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type SignUpForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
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

	decoder := schema.NewDecoder()
	form := SignUpForm{}
	if err := decoder.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}

	fmt.Fprint(w, form)
}
