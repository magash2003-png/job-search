package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	err = service.CreateFavorite(favorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Favorite added"))
}

func GetFavorites(w http.ResponseWriter, r *http.Request) {

	favorites, err := service.GetFavorites()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = service.DeleteFavorite(id)
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