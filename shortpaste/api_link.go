package shortpaste

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/l/"), "/")
	if id == "" {
		var links []Link
		app.db.Find(&links)
		json.NewEncoder(w).Encode(map[string][]Link{"links": links})
	} else {
		var link Link
		if err := app.db.First(&link, "id = ?", id).Error; err == nil {
			http.Redirect(w, r, link.Link, http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Link for `%s` not found!\n", id)
		}
	}
}

func (app *App) handleCreateLink(w http.ResponseWriter, r *http.Request) {
	link := Link{}
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		id := strings.TrimPrefix(r.URL.Path, "/l/")
		if id != "" && link.ID == "" {
			link.ID = id
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": "failed"})
		return
	}
	app.db.Create(&link)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}
