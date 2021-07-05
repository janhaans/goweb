package controllers

import (
	"fmt"
	"net/http"

	"github.com/janhaans/goweb/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
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
	var form SignUpForm
	if err := parseFunc(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Email = %s\n", form.Email)
	fmt.Fprintf(w, "Password = %s\n", form.Password)
}
