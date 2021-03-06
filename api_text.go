package shortpaste

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

func (app *App) handleText(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.handleGetText(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (app *App) handleGetTexts(w http.ResponseWriter, r *http.Request) {
	var texts []Text
	app.db.Find(&texts)
	json.NewEncoder(w).Encode(map[string][]Text{"texts": texts})
}

func (app *App) handleGetText(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/t/"), "/")
	if id == "" {
		onNotFound(w, "No ID found in request")
		return
	}
	var text Text
	if err := app.db.First(&text, "id = ?", id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Text for `%s` not found!\n", id)
		return
	}

	filePath := path.Join(app.storagePath, "data", "texts", text.ID+"."+text.Type)

	if _, ok := r.URL.Query()["download"]; ok {
		w.Header().Set("Content-Disposition", "attachment; filename="+text.ID+"."+text.Type)
		http.ServeFile(w, r, filePath)
		return
	}

	t, err := template.ParseFS(templateFS, "templates/text.html")
	if err != nil {
		onServerError(w, err, "failed to parse template")
		return
	}
	textContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		onServerError(w, err, "failed to read text")
		return
	}
	var highlight string
	if text.NoHighlight {
		highlight = "language-plaintext"
	}
	data := struct {
		ID    string
		Class string
		Text  string
	}{
		ID:    text.ID,
		Class: highlight,
		Text:  html.EscapeString(string(textContent)),
	}
	t.Execute(w, data)
	text.HitCount += 1
	app.db.Save(&text)
}

func (app *App) handleCreateText(w http.ResponseWriter, r *http.Request) {
	text := Text{}
	if err := json.NewDecoder(r.Body).Decode(&text); err != nil {
		onClientError(w, err, "check the input and try again")
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/l/")
	if id != "" && text.ID == "" {
		text.ID = id
	}
	if err := text.validate(); err != nil {
		onClientError(w, err, "check the input and try again")
		return
	}

	if text.Type == "" {
		text.Type = "txt"
	}

	filePath := path.Join(app.storagePath, "data", "texts", text.ID+"."+text.Type)
	if err := os.MkdirAll(path.Dir(filePath), 0700); err != nil {
		onServerError(w, err, "failed to create folder")
		return
	}

	if err := ioutil.WriteFile(filePath, []byte(text.Text), 0600); err != nil {
		onServerError(w, err, "failed to copy contents to disk")
		return
	}
	if err := app.db.Create(&text).Error; err != nil {
		onServerError(w, err, "failed to create DB entry")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}
