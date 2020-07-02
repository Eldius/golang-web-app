package server

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("./templates/*.html"))

/*
IndexHandler index handler
*/
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}

/*
StartServer starts the server
*/
func StartServer(appPort int) {
	http.HandleFunc("/", IndexHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil)
}
