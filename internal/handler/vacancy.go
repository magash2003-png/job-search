package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"job-search/internal/middleware"
	"job-search/internal/models"
	"job-search/internal/service"
)

func CreateVacancy(w http.ResponseWriter, r *http.Request) {

	var vacancy models.Vacancy

	err := json.NewDecoder(r.Body).Decode(&vacancy)
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

	if role != "admin" && role != "employer" {
		http.Error(w, "Only admin and employer can create vacancies", http.StatusForbidden)
		return
	}

	vacancy.EmployerID = userID

	err = service.CreateVacancy(vacancy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Vacancy created successfully"))
}

func GetVacancies(w http.ResponseWriter, r *http.Request) {

	vacancies, err := service.GetVacancies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(vacancies)
}

func Vacancies(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetVacancies(w, r)

	case http.MethodPost:
		CreateVacancy(w, r)

	default:
		http.Error(w, "Method NotAllowed", http.StatusMethodNotAllowed)
	}
}

func Vacancy(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	vacancy, err := service.GetVacancyByID(id)
	if err != nil {
		http.Error(w, "Vacancy not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(vacancy)
}
