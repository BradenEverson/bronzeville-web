package main

import (
	"html/template"
	"net/http"
)

//Return main index file
func mainReq(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("web/index.html"))
    tmpl.Execute(w, nil)
}
//Return static styles css file
func getStyles(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("web/styles/main.css"))
    tmpl.Execute(w, nil)
}
