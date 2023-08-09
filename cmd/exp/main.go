package main

import (
	// "fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta struct {
		Visits int
	}
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Zwe Nyan Zaw",
		Age:  111,
		Meta: struct {
			Visits int
		}{
			Visits: 50,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
