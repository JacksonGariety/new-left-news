package utils

import (
	"html/template"
	"os"
	"path"
	"path/filepath"
	"log"
	"errors"
	"time"
  timeago "github.com/ararog/timeago"
)

var BasePath = os.Getenv("base_path")

var funcMap = template.FuncMap{
	"dict": func (values ...interface{}) (map[string]interface{}, error) {
		if len(values)%2 != 0 {
			return nil, errors.New("invalid dict call")
		}
		dict := make(map[string]interface{}, len(values)/2)
		for i := 0; i < len(values); i+=2 {
			key, ok := values[i].(string)
			if !ok {
				return nil, errors.New("dict keys must be strings")
			}
			dict[key] = values[i+1]
		}
		return dict, nil
	},
	"timeAgoInWords": func (date time.Time) string {
		rel, _ := timeago.TimeAgoWithTime(time.Now().Local().Add(time.Hour * -8), date)
		return rel
	},
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
