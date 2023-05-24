package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.cfg.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
