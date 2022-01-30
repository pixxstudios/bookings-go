package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)

	if err != nil {
		fmt.Println("error parsing template")
		return
	}

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template")
		return
	}
}
