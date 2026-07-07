package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"job-search/internal/middleware"
	"job-search/internal/models"
	"job-search/internal/service"
)

func CreateResume(w http.ResponseWriter, r *http.Request) {

	var resume models.Resume

	err := json.NewDecoder(r.Body).Decode(&resume)
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
			"error": "Only applicants can create resumes",
		})
		return
	}

	resume.UserID = userID

	err = service.CreateResume(resume)
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
		"message": "Resume created successfully",
	})
}

func GetResumes(w http.ResponseWriter, r *http.Request) {

	role, ok := r.Context().Value(middleware.RoleKey).(string)
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

		userID, ok := r.Context().Value(middleware.UserIDKey).(int)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		resumes, err := service.GetResumesByUser(userID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if len(resumes) == 0 {
			json.NewEncoder(w).Encode(map[string]any{
				"message": "No resumes found",
				"data":    []models.Resume{},
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Resumes retrieved successfully",
			"data":    resumes,
		})

	case "employer":

		userID, ok := r.Context().Value(middleware.UserIDKey).(int)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		resumes, err := service.GetResumesForEmployer(userID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if len(resumes) == 0 {
			json.NewEncoder(w).Encode(map[string]any{
				"message": "No resumes found",
				"data":    []models.Resume{},
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Resumes retrieved successfully",
			"data":    resumes,
		})

	case "admin":

		resumes, err := service.GetResumes()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if len(resumes) == 0 {
			json.NewEncoder(w).Encode(map[string]any{
				"message": "No resumes found",
				"data":    []models.Resume{},
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Resumes retrieved successfully",
			"data":    resumes,
		})

	default:

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Access denied",
		})
	}
}

func UpdateResume(w http.ResponseWriter, r *http.Request) {

	var resume models.Resume

	err := json.NewDecoder(r.Body).Decode(&resume)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON",
		})
		return
	}

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID",
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

	if role != "applicant" && role != "admin" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Access denied",
		})
		return
	}

	resume.ID = id

	if role == "admin" {
		err = service.UpdateResumeByAdmin(resume)
	} else {
		resume.UserID = userID
		err = service.UpdateResume(resume)
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Resume updated successfully",
	})
}

func DeleteResume(w http.ResponseWriter, r *http.Request) {

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID",
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

	if role != "applicant" && role != "admin" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Access denied",
		})
		return
	}

	if role == "admin" {
		err = service.DeleteResumeByAdmin(id)
	} else {
		err = service.DeleteResume(id, userID)
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Resume deleted successfully",
	})
}

func Resumes(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetResumes(w, r)

	case http.MethodPost:
		CreateResume(w, r)

	case http.MethodPut:
		UpdateResume(w, r)

	case http.MethodDelete:
		DeleteResume(w, r)

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})
	}
}
