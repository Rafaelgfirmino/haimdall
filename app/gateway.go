package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const PathServiceMap string = "./servicesMap"

var gateway Gateway

type Gateway struct {
	Services []Service
}
type Service struct {
	Name     string    `json:name`
	Url      string    `json:url`
	Handlers []Handler `json:handler`
}
type Handler struct {
	Listen         string `json:listen`
	ContentType    string `json:listem`
	Authorization  bool   `json:authorization`
	ServicePath    string `json:servicePath`
	ServiceFullURL string
}

func init() {
	readAllFilesServices()
}

func readAllFilesServices() {
	files, err := ioutil.ReadDir(PathServiceMap)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Mode().IsRegular() {
			if filepath.Ext(f.Name()) == ".json" {
				gateway.readFileService(f.Name())
			}
		}
	}
}

func (gateway *Gateway) readFileService(fileName string) {
	fileDir := fmt.Sprintf("%s/%s", PathServiceMap, fileName)

	file, e := ioutil.ReadFile(fileDir)
	if e != nil {
		fmt.Printf("File  error: %v\n", e)
		os.Exit(1)
	}
	var tempService Gateway
	json.Unmarshal(file, &tempService)
	gateway.addServicesInGateway(&tempService.Services)
}

func (gateway *Gateway) addServicesInGateway(services *[]Service) {
	for _, service := range *services {
		for keyHandler, handler := range service.Handlers {
			if handler.ServiceFullURL == "" {
				fmt.Println(handler.ServicePath)
				service.Handlers[keyHandler].ServiceFullURL = fmt.Sprintf("%s%s", service.Url, handler.ServicePath)
			}
		}
		gateway.Services = append(gateway.Services, service)
	}
	fmt.Println(gateway)
}
