package utils

import (
	"html/template"
	"os"
	"path"
	"path/filepath"
	"log"
)

var BasePath = os.Getenv("base_path")

func add(a int, b int) int {
	return a + b
}

var funcMap = template.FuncMap{
	"add": add,
}

var templates map[string]*template.Template

// Load templates on program initialisation
func InitTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	tmpls, err := filepath.Glob(path.Join(BasePath, "app/views/*.html"))
	if err != nil {
		log.Fatal(err)
	}

	partials, err := filepath.Glob(path.Join(BasePath, "app/views/partials/*.html"))
	if err != nil {
		log.Fatal(err)
	}

	for _, tmpl := range tmpls {
		files := append(partials, path.Join(BasePath, "app/views/layout.html"))
		files = append(files, tmpl)
		templates[filepath.Base(tmpl)] = template.Must(template.New("base").Funcs(funcMap).ParseFiles(files...))
	}
}
