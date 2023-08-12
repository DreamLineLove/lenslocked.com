package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/DreamLineLove/lenslocked/controllers"
	"github.com/DreamLineLove/lenslocked/templates"
	"github.com/DreamLineLove/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

type faq struct {
	Question string
	Answer   template.HTML
}

func main() {
	r := chi.NewRouter()

	data := []faq{
		{
			"Q: Is there a free version?",
			"A: Yes! We offer a free trial for 30 day",
		},
		{
			"Q: What are your support hours?",
			"A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
		},
		{
			"Q: How do I contact support?",
			`A: Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com`,
		},
	}
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml")), nil))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml")), nil))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "faq.gohtml")), data))

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
