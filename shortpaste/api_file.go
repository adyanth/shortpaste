package shortpaste

import "net/http"

func (app *App) handleFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.handleGetFile(w, r)
	case "POST", "PUT":
		app.handleCreateFile(w, r)
	}
}

func (app *App) handleGetFile(w http.ResponseWriter, r *http.Request) {

}

func (app *App) handleCreateFile(w http.ResponseWriter, r *http.Request) {

}
