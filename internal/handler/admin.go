package handler

import (
	"encoding/json"
	"net/http"

	"job-search/internal/dto"
	"job-search/internal/middleware"
	"job-search/internal/service"
)

func ChangeUserRole(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	role, ok := r.Context().Value(middleware.RoleKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if role != "admin" {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	var request dto.ChangeRoleRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = service.ChangeUserRole(request.UserID, request.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User role updated successfully"))
}
