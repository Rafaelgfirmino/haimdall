package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Olá Mundo")
	OpenFactory()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
