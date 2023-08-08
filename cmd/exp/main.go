package main

import (
	// "fmt"
	"html/template"
	"os"
)

type data struct {
	Language string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Data []data
	}{
		Data: []data{
			data{"German"},
			data{"English"},
		},
	}

	t.Execute(os.Stdout, data)
}
