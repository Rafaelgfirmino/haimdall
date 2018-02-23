package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"strings"
	"github.com/go-fsnotify/fsnotify"
)

var i = 0

func main() {
	Start()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				Start()
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(PathServiceMap); err != nil {
		fmt.Println("ERROR", err)
	}
	OpenPemFactory()
	heimdallPort := fmt.Sprintf(":%s", os.Getenv("HEIMDALL_PORT"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HttpHandler(w, r)
	})

	http.ListenAndServe(heimdallPort, nil)
	<-done

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
