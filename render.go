package main

import (
	"fmt"
	"net/http"
	"text/template"
)

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
