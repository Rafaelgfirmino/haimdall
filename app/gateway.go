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
	Listem        string `json:listem`
	ContentType   string `json:listem`
	Authorization bool   `json:authorization`
	ServicePath   string `json:servicePath`
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
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var tempService Gateway
	json.Unmarshal(file, &tempService)
	fmt.Println(tempService)
	// gateway.addServicesInGateway(&tempService)
}

func (gateway *Gateway) addServicesInGateway(services *[]Service) {
	for _, service := range *services {
		gateway.Services = append(gateway.Services, service)
	}
}
