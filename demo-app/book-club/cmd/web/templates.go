package main

import (
	"fast-track-istio/tour-guide/pkg/models"
	"html/template"
	"path/filepath"
	"strings"
	"time"
)

type requestMetadata struct {
	Source          string
	InvokedBy       string
	RequestDuration time.Duration
}

type booksTemplateData struct {
	RequestMetadata *requestMetadata
	Book            *models.Book
	Books           *[]*models.Book
}

type moviesTemplateData struct {
	RequestMetadata *requestMetadata
	Movie           *models.Movie
	Movies          *[]*models.Movie
}

type templateData struct {
	Error              []string
	BooksTemplateData  *booksTemplateData
	MoviesTemplateData *moviesTemplateData
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format("02 Jan 2006")
}

func titleCase(str string) string {
	return strings.Title(strings.ToLower(str))
}

var functions = template.FuncMap{
	"humanDate": humanDate,
	"titleCase": titleCase,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
