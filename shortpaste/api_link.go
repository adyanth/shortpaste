package shortpaste

import (
	"net/http"
)

func (app *App) handleLink(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.handleGetLink(w, r)
	case "POST", "PUT":
		app.handleCreateLink(w, r)
	}
}

func (app *App) handleGetLink(w http.ResponseWriter, r *http.Request) {

}

func (app *App) handleCreateLink(w http.ResponseWriter, r *http.Request) {

}
