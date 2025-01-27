package utils

import (
	"awesomeProject/internal/email/models"
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, detailsMessage ...string) {
	details := ""
	if len(detailsMessage) > 0 {
		details = detailsMessage[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.ErrorResponseType{
		Message:      message,
		StatusCode:   statusCode,
		Error:        true,
		DetailsError: details,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
