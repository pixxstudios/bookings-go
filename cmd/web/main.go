package main

import (
	"fmt"
	"net/http"

	"github.com/pixxstudios/bookings-go/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Application started on %s", portNumber))

	http.ListenAndServe(portNumber, nil)
}
