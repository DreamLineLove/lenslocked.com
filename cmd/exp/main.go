package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "os"
)

type User struct {
	Name   string
	Text   string
	Html   template.HTML
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:   "Zwe Nyan Zaw",
		Text:   `<script>alert("This is a very nefarious attack");</script>`,
		Html:   `<script>alert("This is expected to run");</script>`,
		Visits: 100,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html, charset=utf-8")
		err = t.Execute(w, user)
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Starting the server on :3030")
	err = http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}
