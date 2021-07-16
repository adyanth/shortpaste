package shortpaste

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
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
		filePath := path.Join(app.storagePath, "files", file.ID, file.Name)
		if _, ok := r.URL.Query()["download"]; ok {
			w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
			http.ServeFile(w, r, filePath)
			return
		}

		t, err := template.ParseFiles("static/file.html")
		if err != nil {
			onServerError(w, err, "failed to parse template")
			return
		}
		fi, err := os.Stat(filePath)
		if err != nil {
			onServerError(w, err, "failed to get file size")
			return
		}
		data := struct {
			Name  string
			Src   string
			Image bool
			Size  string
		}{
			Name:  file.Name,
			Src:   "/f/" + id + "?download",
			Image: strings.HasPrefix(file.MIME, "image/"),
			Size:  IECFormat(fi.Size()),
		}
		t.Execute(w, data)
	}
}

func (app *App) handleCreateFile(w http.ResponseWriter, r *http.Request) {
	file := File{}
	file.ID = strings.TrimPrefix(r.URL.Path, "/f/")
	if err := file.validate(); err != nil {
		onClientError(w, err, "check the input and try again")
		return
	}

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		onClientError(w, err, "failed to retrieve file, check if the upload completed")
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
	if err := os.MkdirAll(path.Dir(filePath), 0700); err != nil {
		onServerError(w, err, "failed to create folder")
		return
	}

	dst, err := os.Create(filePath)
	if err != nil {
		onServerError(w, err, "failed to create file handler")
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, uploadedFile); err != nil {
		onServerError(w, err, "failed to copy contents to disk")
		return
	}
	if err = app.db.Create(&file).Error; err != nil {
		onServerError(w, err, "failed to create DB entry")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}
