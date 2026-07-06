package handler

import (
	"encoding/json"
	"net/http"

	"job-search/internal/middleware"
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

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	role, ok := r.Context().Value(middleware.RoleKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if role != "applicant" {
		http.Error(w, "Only applicants can create responses", http.StatusForbidden)
		return
	}

	response.UserID = userID

	err = service.CreateResponse(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Response created"))
}

func GetResponses(w http.ResponseWriter, r *http.Request) {

	responses, err := service.GetResponses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(responses)
}

func Responses(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetResponses(w, r)

	case http.MethodPost:
		CreateResponse(w, r)

	default:
		http.Error(w, "Method NotAllowed", http.StatusMethodNotAllowed)
	}
}
