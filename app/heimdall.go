package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var i = 0

func main() {

	OpenPemFactory()
	heimdallPort := fmt.Sprintf(":%s", os.Getenv("HEIMDALL_PORT"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HttpHandler(w, r)
	})

	http.ListenAndServe(heimdallPort, nil)
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	originalPathRequest := r.URL.Path
	var handler Handler

	for _, service := range gateway.Services {
		for _, gatewayhandler := range service.Handlers {

			if gatewayhandler.Listen == originalPathRequest {
				handler = gatewayhandler
			}
		}
	}
	if handler.Listen == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}
	var req *http.Request
	fmt.Println(handler)
	req = redirectRequestToService(r, handler)
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)

	//resp with original Content-Type
	headerResp := strings.Join(resp.Header["Content-Type"], "")
	w.Header().Set("Content-Type", headerResp)
	w.Write([]byte(body))
	// fmt.Fprintf(w, string(body))
}

func redirectRequestToService(r *http.Request, handler Handler) *http.Request {
	fmt.Println(handler.ServiceFullURL)
	req, err := http.NewRequest(r.Method, handler.ServiceFullURL, r.Body)
	CheckErr(err)
	req.Header.Set("Content-Type", handler.ContentType)
	return req
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
