package main

import (
	"fmt"
	"net/http"
	"os"
)

var i = 0

func main() {

	OpenFactory()
	heimdallPort := fmt.Sprintf(":%s", os.Getenv("HEIMDALL_PORT"))

	http.HandleFunc("/", handler)
	http.ListenAndServe(heimdallPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	i += 1

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(fmt.Sprintf("%d", i)))

}
