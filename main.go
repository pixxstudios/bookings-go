package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func main() {
	http.HandleFunc("/", Home)

	fmt.Println(fmt.Sprintf("Application started on %s", portNumber))

	http.ListenAndServe(portNumber, nil)
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
