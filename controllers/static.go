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
