package controllers

import (
	"net/http"

	"github.com/DreamLineLove/lenslocked/views"
)

func StaticHandler(tpl views.Template, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, data)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
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

}
