package main

import (
	"github.com/ihaveint/snippetbox/pkg/models"
	"html/template"
	"path/filepath"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		template, err := template.ParseFiles(page)

		if err != nil {
			return nil, err
		}

		template, err = template.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))

		if err != nil {
			return nil, err
		}

		template, err = template.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err

		}
		cache[name] = template
	}
	return cache, nil
}
