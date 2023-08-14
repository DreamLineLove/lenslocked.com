package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Email address: %v\n", r.FormValue("email_address"))
	fmt.Fprintf(w, "Password: %v\n", r.FormValue("password"))
	fmt.Fprintf(w, "Something: %v\n", r.FormValue("something"))
}
