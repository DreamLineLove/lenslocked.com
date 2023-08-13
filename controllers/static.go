package controllers

import (
	"html/template"
	"net/http"

	"github.com/DreamLineLove/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Q: Is there a free version?",
			Answer:   "A: Yes! We offer a free trial for 30 day",
		},
		{
			Question: "Q: What are your support hours?",
			Answer:   "A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
		},
		{
			Question: "Q: How do I contact support?",
			Answer:   `A: Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
