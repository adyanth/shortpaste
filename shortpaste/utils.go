package shortpaste

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

func (link *Link) validate() error {
	return validator.New().Struct(link)
}

func (file *File) validate() error {
	return validator.New().Struct(file)
}

func (text *Text) validate() error {
	return validator.New().Struct(text)
}

func onClientError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": msg})
}

func onServerError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err), "message": msg})
}
