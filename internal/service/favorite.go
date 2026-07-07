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

	exists, err := repository.VacancyExists(f.VacancyID)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("vacancy not found")
	}

	exists, err = repository.FavoriteExists(f.UserID, f.VacancyID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("vacancy already in favorites")
	}

	return repository.CreateFavorite(f)
}

func GetFavorites(userID int) ([]models.Favorite, error) {

	return repository.GetFavorites(userID)

}

func DeleteFavorite(id int, userID int) error {

	if id <= 0 {
		return errors.New("invalid favorite id")
	}

	return repository.DeleteFavorite(id, userID)
}
