package main

import (
	"context"
	"gh_static_portfolio/cmd/templates"
	"log"
	"os"
)

func main() {
	// This is a comment
	log.Println("Hello, World!")
	f, err := os.Create("./docs/hello.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = templates.TestComponent().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
