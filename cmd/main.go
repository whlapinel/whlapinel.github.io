package main

import (
	"gh_static_portfolio/cmd/templates"
	"log"
)

func main() {
	err := templates.RenderPages()
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	log.Println("pages rendered")
}
