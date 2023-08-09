package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")

	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Happened when parsing template files: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Happened when executing the template: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, `<html>
		<h1>Frequently Asked Questions</h1>
  	<ul>
    	<li>
      	<h2>Q: Is there a free version?</h2>
    		<p>A: Yes! We offer a free trial for 30 day</p>
    	</li>
    	<li>
      	<h2>Q: What are your support hours?</h2>
    		<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p>
    	</li>
    	<li>
      	<h2>Q: How do I contact support?</h2>
    		<p>A: Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</p>
    	</li>
  	</ul>
  </html>
	`)
}

func notFoundHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>404 page not found</h1>")
}

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandlerFunc)
	r.Get("/contact", contactHandlerFunc)
	r.Get("/faq", faqHandlerFunc)
	r.NotFound(notFoundHandlerFunc)

	fmt.Fprintln(os.Stdout, "Starting the server on :3000...")

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
