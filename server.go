package shortpaste

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (app *App) handleRequests() {
	// Short links
	http.HandleFunc("/f/", app.handleFile)
	http.HandleFunc("/l/", app.handleLink)
	http.HandleFunc("/t/", app.handleText)
	// Admin API
	http.HandleFunc("/api/v1/", app.basicAuth(app.handleAPI))
	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", app.basicAuth(fs.ServeHTTP))

	fmt.Printf("Server starting at %s\n", app.bind)
	log.Fatal(http.ListenAndServe(app.bind, nil))
}

func (app *App) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.noAuth {
			username, password, ok := r.BasicAuth()
			if ok {
				// Calculate SHA-256 hashes for the provided and expected
				// usernames and passwords.
				usernameHash := sha256.Sum256([]byte(username))
				passwordHash := sha256.Sum256([]byte(password))
				expectedUsernameHash := sha256.Sum256([]byte(app.username))
				expectedPasswordHash := sha256.Sum256([]byte(app.password))

				usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
				passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

				if usernameMatch && passwordMatch {
					next.ServeHTTP(w, r)
					return
				}
			}

			// If the Authentication header is not present, is invalid, or the
			// username or password is wrong, then set a WWW-Authenticate
			// header to inform the client that we expect them to use basic
			// authentication and send a 401 Unauthorized response.
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			// If noAuth is set, bypass auth
			next.ServeHTTP(w, r)
			return
		}
	})
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
