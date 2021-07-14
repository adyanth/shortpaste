package shortpaste

import (
	"fmt"
	"log"
	"net/http"
)

func (app *App) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (app *App) handleRequests() {
	http.HandleFunc("/", app.homePage)
	http.HandleFunc("/f/", app.handleFile)
	http.HandleFunc("/l/", app.handleLink)
	http.HandleFunc("/t/", app.handleText)
	fmt.Printf("Server starting at %s\n", app.bind)
	log.Fatal(http.ListenAndServe(app.bind, nil))
}
