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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON",
		})
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Unauthorized",
		})
		return
	}

	role, ok := r.Context().Value(middleware.RoleKey).(string)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Unauthorized",
		})
		return
	}

	if role != "applicant" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Only applicants can create responses",
		})
		return
	}

	response.UserID = userID

	err = service.CreateResponse(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Response created successfully",
	})
}

func GetResponses(w http.ResponseWriter, r *http.Request) {

	role, ok := r.Context().Value(middleware.RoleKey).(string)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Unauthorized",
		})
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Unauthorized",
		})
		return
	}

	switch role {

	case "applicant":

		responses, err := service.GetResponsesByUser(userID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if len(responses) == 0 {
			json.NewEncoder(w).Encode(map[string]any{
				"message": "No responses found",
				"data":    []models.Response{},
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Responses retrieved successfully",
			"data":    responses,
		})

	case "employer":

		responses, err := service.GetResponsesByEmployer(userID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if len(responses) == 0 {
			json.NewEncoder(w).Encode(map[string]any{
				"message": "No responses found",
				"data":    []models.Response{},
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Responses retrieved successfully",
			"data":    responses,
		})

	case "admin":

		responses, err := service.GetResponses()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if len(responses) == 0 {
			json.NewEncoder(w).Encode(map[string]any{
				"message": "No responses found",
				"data":    []models.Response{},
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Responses retrieved successfully",
			"data":    responses,
		})

	default:

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Access denied",
		})
	}
}

func Responses(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetResponses(w, r)

	case http.MethodPost:
		CreateResponse(w, r)

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})
	}
}
