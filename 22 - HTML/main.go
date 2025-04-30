package main

import (
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{Nome: "Lucas", Email: "lucas.ribeiro@gmail.com"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	print("Server started at :8080\n")
	http.ListenAndServe(":8080", nil)
}
