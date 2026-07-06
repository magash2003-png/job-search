package service

import (
	"errors"

	"job-search/internal/models"
	"job-search/internal/repository"
)

func CreateFavorite(f models.Favorite) error {

	if f.UserID == 0 {
		return errors.New("user_id is required")
	}

	if f.VacancyID == 0 {
		return errors.New("vacancy_id is required")
	}

	return repository.CreateFavorite(f)
}

func GetFavorites() ([]models.Favorite, error) {

	return repository.GetFavorites()

}

func DeleteFavorite(id int, userID int) error {

	if id <= 0 {
		return errors.New("invalid favorite id")
	}

	return repository.DeleteFavorite(id, userID)
}
