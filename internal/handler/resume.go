package handler

import (
	"encoding/json"
	"net/http"

	"job-search/internal/middleware"
	"job-search/internal/models"
	"job-search/internal/service"
)

func CreateResume(w http.ResponseWriter, r *http.Request) {

	var resume models.Resume

	err := json.NewDecoder(r.Body).Decode(&resume)
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
		http.Error(w, "Only applicants can create resumes", http.StatusForbidden)
		return
	}

	resume.UserID = userID

	err = service.CreateResume(resume)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Resume created"))
}

func GetResumes(w http.ResponseWriter, r *http.Request) {

	resumes, err := service.GetResumes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resumes)
}

func Resumes(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetResumes(w, r)

	case http.MethodPost:
		CreateResume(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
