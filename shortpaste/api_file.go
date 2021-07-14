package shortpaste

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func (app *App) handleFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.handleGetFile(w, r)
	case "POST", "PUT":
		app.handleCreateFile(w, r)
	}
}

func (app *App) handleGetFile(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/f/"), "/")
	if id == "" {
		var files []File
		app.db.Find(&files)
		json.NewEncoder(w).Encode(map[string][]File{"files": files})
	} else {
		var file File
		if err := app.db.First(&file, "id = ?", id).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Link for `%s` not found!\n", id)
			return
		}
		if _, ok := r.URL.Query()["download"]; ok || !strings.HasPrefix(file.MIME, "image/") {
			w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
		}
		filePath := path.Join(app.storagePath, "files", file.ID, file.Name)
		http.ServeFile(w, r, filePath)
	}
}

func (app *App) handleCreateFile(w http.ResponseWriter, r *http.Request) {
	file := File{}
	file.ID = strings.TrimPrefix(r.URL.Path, "/f/")
	if file.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "No ID provided!", "message": "failed to retrieve id"})
		return
	}

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": "failed to retrieve file"})
		return
	}
	defer uploadedFile.Close()

	file.Name = handler.Filename
	file.MIME = handler.Header["Content-Type"][0]

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	filePath := path.Join(app.storagePath, "files", file.ID, file.Name)

	// Create folder and file
	err = os.MkdirAll(path.Dir(filePath), 0700)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": "failed to create folder"})
		return
	}

	dst, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": "failed to create file handler"})
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, uploadedFile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": "failed to copy contents to disk"})
		return
	}
	if err = app.db.Create(&file).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": "failed to create DB entry"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}
