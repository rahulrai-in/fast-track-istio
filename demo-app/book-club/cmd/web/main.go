package main

import (
	"flag"
	"html/template"
	"net/http"
	"os"
)

type application struct {
	templateCache map[string]*template.Template
	settings      *settings
}

type settings struct {
	booksServiceUri       string
	moviesServiceUri      string
	independentServiceUri string
}

func main() {
	// port settings
	addr := flag.String("addr", ":4000", "Server port")
	flag.Parse()

	// application settings
	settings := &settings{}
	settings.moviesServiceUri = os.Getenv("moviesServiceUri")
	settings.booksServiceUri = os.Getenv("booksServiceUri")
	settings.independentServiceUri = os.Getenv("independentServiceUri")

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		panic(err)
	}

	app := application{
		templateCache: templateCache,
		settings:      settings,
	}

	mux := app.routes()
	srv := http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
