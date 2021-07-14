package shortpaste

import "net/http"

func (app *App) handleText(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.handleGetText(w, r)
	case "POST", "PUT":
		app.handleCreateText(w, r)
	}
}

func (app *App) handleGetText(w http.ResponseWriter, r *http.Request) {

}

func (app *App) handleCreateText(w http.ResponseWriter, r *http.Request) {

}
