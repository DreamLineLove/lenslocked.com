package main

import (
	"fmt"
	"net/http"
	"os"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to Lenslocked!</h1>")
}

// Outcome
// - when /contact is requested by the browser, our server should handle it
// - the server should reply with an html and the browser should know that it is indeed html
// - the html should include an a tag with the correct href

func contactHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:zwenyanzaw@protonmail.com\">zwenyanzaw@protonmail.com</a>.")
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
	// fmt.Fprint(w, "<h1>Q: Is there a free version?</h1><p>A: Yes! We offer a free trial for 30 day</h1>s on any paid plans.</p>")
	// fmt.Fprint(w, "<h1>Q: What are your support hours?</h1><p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p>")
	// fmt.Fprint(w, "<h1>Q: How do I contact support?</h1><p>A: Email us - <a href=\"mailto:support@lenslocked.com\">support@Lenslocked.com</a></p>")
}

func notFoundHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>404 page not found</h1>")
}

// We want to write code in our server that returns the url path and url raw path as part of the response body to whatever path

// func generalHandlerFunc(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandlerFunc(w, r)
// 	case "/contact":
// 		contactHandlerFunc(w, r)
// 	default:
// 		http.NotFound(w, r)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandlerFunc(w, r)
	case "/contact":
		contactHandlerFunc(w, r)
	case "/faq":
		faqHandlerFunc(w, r)
	default:
		// http.NotFound(w, r)
		notFoundHandlerFunc(w, r)
	}
}

func main() {
	router := Router{}

	fmt.Fprintln(os.Stdout, "Starting the server on :3000...")

	if err := http.ListenAndServe(":3000", router); err != nil {
		panic(err)
	}
}
