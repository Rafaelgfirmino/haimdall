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

type Gateway struct {
	Services []Service
}
type Service struct {
	Name string `json`
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
				readFileService(f.Name())
			}
		}
	}
}

func readFileService(fileName string) {
	fileDir := fmt.Sprintf("%s/%s", PathServiceMap, fileName)

	file, e := ioutil.ReadFile(fileDir)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var tempService []Service
	json.Unmarshal(file, &tempService)
}
