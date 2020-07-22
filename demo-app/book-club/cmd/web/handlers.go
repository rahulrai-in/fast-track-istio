package main

import (
	"encoding/json"
	"fast-track-istio/tour-guide/pkg/models"
	"io"
	"net/http"
	"sync"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl", nil)
}

func (app *application) movies(w http.ResponseWriter, r *http.Request) {
	var (
		moviesResponse        []*models.Movie
		moviesRequestDuration time.Duration
	)

	templateData := &templateData{
		MoviesTemplateData: &moviesTemplateData{
			RequestMetadata: &requestMetadata{
				Source:          app.settings.moviesServiceUri,
				InvokedBy:       "movies",
				RequestDuration: moviesRequestDuration,
			},
			Movies: &moviesResponse,
		},
	}

	err, bodyStream, duration := app.httpGet(app.settings.moviesServiceUri + "/movies")
	defer bodyStream.Close()
	templateData.MoviesTemplateData.RequestMetadata.RequestDuration = duration
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	err = json.NewDecoder(bodyStream).Decode(&moviesResponse)
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	app.render(w, r, "movies.page.tmpl", templateData)
}

func (app *application) books(w http.ResponseWriter, r *http.Request) {
	var booksResponse []*models.Book
	response := &booksTemplateData{
		RequestMetadata: &requestMetadata{
			Source:    app.settings.booksServiceUri,
			InvokedBy: "books",
		},
		Books: &booksResponse,
	}

	templateData := &templateData{
		BooksTemplateData: response,
	}

	err, bodyStream, duration := app.httpGet(app.settings.booksServiceUri + "/books")
	defer bodyStream.Close()
	response.RequestMetadata.RequestDuration = duration
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	err = json.NewDecoder(bodyStream).Decode(&booksResponse)
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	app.render(w, r, "books.page.tmpl", templateData)
}

func (app *application) independent(w http.ResponseWriter, r *http.Request) {
	var (
		booksResponse         []*models.Book
		moviesResponse        []*models.Movie
		booksRequestDuration  time.Duration
		moviesRequestDuration time.Duration
		wg                    sync.WaitGroup
	)

	templateData := &templateData{
		BooksTemplateData: &booksTemplateData{
			RequestMetadata: &requestMetadata{
				Source:          app.settings.independentServiceUri,
				InvokedBy:       "independent",
				RequestDuration: booksRequestDuration,
			},
			Books: &booksResponse,
		},
		MoviesTemplateData: &moviesTemplateData{
			RequestMetadata: &requestMetadata{
				Source:          app.settings.independentServiceUri,
				InvokedBy:       "independent",
				RequestDuration: moviesRequestDuration,
			},
			Movies: &moviesResponse,
		},
	}

	wg.Add(2)
	go func() {
		err, bodyStream, duration := app.httpGet(app.settings.independentServiceUri + "/books")
		defer bodyStream.Close()
		defer wg.Done()
		templateData.BooksTemplateData.RequestMetadata.RequestDuration = duration
		if err != nil {
			templateData.Error = append(templateData.Error, err.Error())
		}

		err = json.NewDecoder(bodyStream).Decode(&booksResponse)
		if err != nil {
			templateData.Error = append(templateData.Error, err.Error())
		}
	}()

	go func() {
		err, bodyStream, duration := app.httpGet(app.settings.independentServiceUri + "/movies")
		defer bodyStream.Close()
		defer wg.Done()
		templateData.MoviesTemplateData.RequestMetadata.RequestDuration = duration
		if err != nil {
			templateData.Error = append(templateData.Error, err.Error())
		}

		err = json.NewDecoder(bodyStream).Decode(&moviesResponse)
		if err != nil {
			templateData.Error = append(templateData.Error, err.Error())
		}
	}()

	wg.Wait()
	app.render(w, r, "independent.page.tmpl", templateData)
}

func (app *application) legacy(w http.ResponseWriter, r *http.Request) {
	var booksResponse []*models.Book
	response := &booksTemplateData{
		RequestMetadata: &requestMetadata{
			Source:    app.settings.booksServiceUri,
			InvokedBy: "legacy",
		},
		Books: &booksResponse,
	}

	templateData := &templateData{
		BooksTemplateData: response,
	}

	err, bodyStream, duration := app.httpGet(app.settings.booksServiceUri + "/books/legacy")
	defer bodyStream.Close()
	response.RequestMetadata.RequestDuration = duration
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	err = json.NewDecoder(bodyStream).Decode(&booksResponse)
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	app.render(w, r, "legacy.page.tmpl", templateData)
}

func (app *application) rentBook(w http.ResponseWriter, r *http.Request) {
	var (
		bookResponse models.Book
		err          error
		bodyStream   io.ReadCloser
		duration     time.Duration
	)
	response := &booksTemplateData{
		RequestMetadata: &requestMetadata{
			Source: app.settings.booksServiceUri,
		},
		Book: &bookResponse,
	}

	templateData := &templateData{
		BooksTemplateData: response,
	}

	switch requester := r.URL.Query().Get("src"); requester {
	case "books":
		err, bodyStream, duration = app.httpGet(app.settings.booksServiceUri + "/books/" + r.URL.Query().Get("id"))
	case "independent":
		err, bodyStream, duration = app.httpGet(app.settings.independentServiceUri + "/books/" + r.URL.Query().Get("id"))
	}

	defer bodyStream.Close()
	response.RequestMetadata.RequestDuration = duration
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	err = json.NewDecoder(bodyStream).Decode(&bookResponse)
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	app.render(w, r, "rent-book.page.tmpl", templateData)
}

func (app *application) rentMovie(w http.ResponseWriter, r *http.Request) {
	var (
		movieResponse models.Movie
		err           error
		bodyStream    io.ReadCloser
		duration      time.Duration
	)

	templateData := &templateData{
		MoviesTemplateData: &moviesTemplateData{
			RequestMetadata: &requestMetadata{
				Source: app.settings.booksServiceUri,
			},
			Movie: &movieResponse,
		},
	}

	switch requester := r.URL.Query().Get("src"); requester {
	case "movies":
		err, bodyStream, duration = app.httpGet(app.settings.moviesServiceUri + "/movies/" + r.URL.Query().Get("id"))
	case "independent":
		templateData.MoviesTemplateData.RequestMetadata.Source = app.settings.independentServiceUri
		err, bodyStream, duration = app.httpGet(app.settings.independentServiceUri + "/movies/" + r.URL.Query().Get("id"))
	}

	defer bodyStream.Close()
	templateData.MoviesTemplateData.RequestMetadata.RequestDuration = duration
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	err = json.NewDecoder(bodyStream).Decode(&movieResponse)
	if err != nil {
		templateData.Error = append(templateData.Error, err.Error())
	}

	app.render(w, r, "rent-movie.page.tmpl", templateData)
}
