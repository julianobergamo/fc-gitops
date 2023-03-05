package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!!!</h1>")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Juliano T. Bergamo <julianobergamo@gmail.com>")
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/contact", Contact)
	http.ListenAndServe(":8080", nil)
}
