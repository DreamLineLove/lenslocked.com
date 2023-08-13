package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DreamLineLove/lenslocked/controllers"
	"github.com/DreamLineLove/lenslocked/templates"
	"github.com/DreamLineLove/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "home.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "contact.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "faq.gohtml"))))

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
