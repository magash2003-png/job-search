package handler

import (
	"encoding/json"
	"net/http"

	"job-search/internal/models"
	"job-search/internal/service"
)

func CreateResponse(w http.ResponseWriter, r *http.Request) {

	var response models.Response

	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = service.CreateResponse(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Response created"))
}

func GetResponses(w http.ResponseWriter, r *http.Request) {

	responses, err := service.GetResponses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(responses)
}

func Responses(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetResponses(w, r)

	case http.MethodPost:
		CreateResponse(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}