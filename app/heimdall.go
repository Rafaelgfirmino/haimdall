package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	OpenFactory()
	heimdallPort := fmt.Sprintf(":%s", os.Getenv("HEIMDALL_PORT"))

	http.HandleFunc("/", handler)
	http.ListenAndServe(heimdallPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
