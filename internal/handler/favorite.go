package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"job-search/internal/middleware"
	"job-search/internal/models"
	"job-search/internal/service"
)

func CreateFavorite(w http.ResponseWriter, r *http.Request) {

	var favorite models.Favorite

	err := json.NewDecoder(r.Body).Decode(&favorite)
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
		http.Error(w, "Only applicants can add favorites", http.StatusForbidden)
		return
	}

	favorite.UserID = userID

	err = service.CreateFavorite(favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Favorite added"))
}

func GetFavorites(w http.ResponseWriter, r *http.Request) {

	favorites, err := service.GetFavorites()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(favorites)
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = service.DeleteFavorite(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Favorite deleted"))
}

func Favorites(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetFavorites(w, r)

	case http.MethodPost:
		CreateFavorite(w, r)

	case http.MethodDelete:
		DeleteFavorite(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
