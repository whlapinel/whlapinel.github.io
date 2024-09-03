package templates

import (
	"strings"

	"github.com/a-h/templ"
)

const (
	rootDir      = "./docs/"
	aboutDir     = "./docs/about/"
	educationDir = "./docs/about/education/"
	blogsDir     = "./docs/blogs/"
)

func DirectoryList() []string {
	return []string{
		rootDir,
		aboutDir,
		educationDir,
		blogsDir,
	}
}

type Templifier interface {
	Templify() templ.Component
	Title() string
	Directory() string
}

type Titler interface {
	Title() string
}

func FileName(t Titler) string {
	if t.Title() == "Home" {
		return "index.html"
	}
	return strings.ReplaceAll(strings.ToLower(t.Title()), " ", "-") + ".html"
}

func RoutePath(directory string) string {
	return strings.ReplaceAll(directory, "./docs", "")
}
