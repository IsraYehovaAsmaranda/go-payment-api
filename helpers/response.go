package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/IsraYehovaAsmaranda/go-payment-api/models"
	"github.com/sirupsen/logrus"
)

func RespondWithError(w http.ResponseWriter, code int, errors string, msg string) {
	if code > 499 {
		logrus.Println("Responding with 5xx error: ", msg)
	}

	err := models.ErrorData{Error: errors}

	RespondWithJSON(w, code, err, msg)
}
func RespondWithJSON(w http.ResponseWriter, code int, payload any, message string) {
	response := models.CommonResponse{
		Status:  code,
		Message: message,
		Data:    payload,
	}

	dat, err := json.Marshal(response)
	if err != nil {
		logrus.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
