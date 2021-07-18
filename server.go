package shortpaste

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (app *App) handleRequests() {
	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	// Short links
	http.HandleFunc("/f/", app.handleFile)
	http.HandleFunc("/l/", app.handleLink)
	http.HandleFunc("/t/", app.handleText)
	// Admin API
	http.HandleFunc("/api/v1/", app.handleAPI)

	fmt.Printf("Server starting at %s\n", app.bind)
	log.Fatal(http.ListenAndServe(app.bind, nil))
}

func (app *App) handleAPI(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/v1")
	switch r.Method {
	case "GET":
		switch {
		case strings.HasPrefix(r.URL.Path, "/f/"):
			app.handleGetFiles(w, r)
		case strings.HasPrefix(r.URL.Path, "/l/"):
			app.handleGetLinks(w, r)
		case strings.HasPrefix(r.URL.Path, "/t/"):
			app.handleGetTexts(w, r)
		default:
			onNotFound(w, "No such endpoint")
		}
	case "POST":
		switch {
		case strings.HasPrefix(r.URL.Path, "/f/"):
			app.handleCreateFile(w, r)
		case strings.HasPrefix(r.URL.Path, "/l/"):
			app.handleCreateLink(w, r)
		case strings.HasPrefix(r.URL.Path, "/t/"):
			app.handleCreateText(w, r)
		default:
			onNotFound(w, "No such endpoint")
		}
	}
}
