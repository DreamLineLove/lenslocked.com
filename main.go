package main

import (
	"fmt"
	// "html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/DreamLineLove/lenslocked/controllers"
	"github.com/DreamLineLove/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html, charset=utf-8")
		fmt.Fprint(w, "<h1>404 page not found</h1>")
	})

	fmt.Fprintln(os.Stdout, "Starting the server on :3000...")

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
