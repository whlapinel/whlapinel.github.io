package main

import (
	"gh_static_portfolio/cmd/data"
	"log"
)

func main() {
	courseInstances, err := data.GenerateInstancesFromCSV()
	if err != nil {
		log.Fatal(err)
	}
	err = data.WriteInstancesToCSV(courseInstances)
	if err != nil {
		log.Fatalf("error writing schedule to csv, %v", err)
	}

}
