package main

import (
	"html/template"
	"net/http"
)

func mainReq(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("web/index.html"))
    tmpl.Execute(w, nil)
}

func getStyles(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("web/styles/main.css"))
    tmpl.Execute(w, nil)
}
