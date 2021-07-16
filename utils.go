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

//IECFormat prints bytes in the International Electrotechnical Commission format
func IECFormat(sizeBytes int64) string {
	suffix := "B"
	num := float64(sizeBytes)
	units := []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}
	for _, unit := range units {
		if num < 1024.0 {
			return fmt.Sprintf("%3.1f%s%s", num, unit, suffix)
		}
		num = (num / 1024)
	}
	return fmt.Sprintf("%.1f%s%s", num, "Yi", suffix)
}
