package main

import (
	"fmt"
	"net/http"
	"os"
)

var i = 0

func main() {

	OpenPemFactory()
	heimdallPort := fmt.Sprintf(":%s", os.Getenv("HEIMDALL_PORT"))

	http.HandleFunc("/teste", handler)
	http.ListenAndServe(heimdallPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i++
	w.Write([]byte(fmt.Sprintf("%d", i)))

}
