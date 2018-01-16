package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func init() {
	files, err := ioutil.ReadDir("./servicesMap")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Mode().IsRegular() {
			if filepath.Ext(f.Name()) == ".json" {
				fmt.Println(f.Name())
			}
		}
	}
}
