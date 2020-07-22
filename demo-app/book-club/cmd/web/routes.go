package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {

	// build routes using pat. Pat will create a mux rather than http.NewServeMux below.
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/books", http.HandlerFunc(app.books))
	mux.Get("/movies", http.HandlerFunc(app.movies))
	mux.Get("/independent", http.HandlerFunc(app.independent))
	mux.Get("/legacy", http.HandlerFunc(app.legacy))
	mux.Get("/rent-book", http.HandlerFunc(app.rentBook))
	mux.Get("/rent-movie", http.HandlerFunc(app.rentMovie))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
