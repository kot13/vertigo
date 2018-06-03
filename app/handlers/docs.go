package handlers

import (
	"net/http"
	"text/template"
)

type Advert struct {
}

func Docs(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/docs.html")

	advert := Advert{}
	t.Execute(w, advert)
}
