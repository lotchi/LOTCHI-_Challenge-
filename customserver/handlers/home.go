package handlers

import (
	"html/template"
	"net/http"
)

// Home handler of index page
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
