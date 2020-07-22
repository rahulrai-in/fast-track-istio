package main

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}

	// write to buffer before sending to client so that any errors in HTML template don't return a partial page.
	buffer := &bytes.Buffer{}
	err := ts.Execute(buffer, td)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = buffer.WriteTo(w)
}

func (app *application) httpGet(url string) (error, io.ReadCloser, time.Duration) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	start := time.Now()
	r, err := httpClient.Get(url)
	if err != nil {
		return err, nil, time.Since(start)
	}

	return nil, r.Body, time.Since(start)
}
