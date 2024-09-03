package templates

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
)

type Templifier interface {
	Templify() templ.Component
	Title() string
	Directory() string
}

func FileName(t Templifier) string {
	if t.Title() != "Home" {
		return "index.html"
	}
	return strings.ReplaceAll(strings.ToLower(t.Title()), " ", "-") + ".html"
}

func RenderPages() error {
	ClearHTMLFiles(AboutDir)
	ClearHTMLFiles(BlogsDir)
	return nil
}

func ClearHTMLFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			err := os.Remove(directory + file.Name())
			if err != nil {
				log.Fatalf("failed to delete file: %v", err)
			}
		}
	}
}

func RenderPage(t Templifier) error {
	f, err := os.Create(t.Directory() + FileName(t))
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	err = t.Templify().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}
