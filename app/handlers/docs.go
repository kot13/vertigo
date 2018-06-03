package handlers

import (
	"net/http"
	"text/template"
)

func Docs(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/docs.html")

	t.Execute(w, struct{}{})
}
