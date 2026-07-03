package service

import (
	"errors"

	"job-search/internal/models"
	"job-search/internal/repository"
)

func CreateResponse(r models.Response) error {

	if r.UserID == 0 {
		return errors.New("user_id is required")
	}

	if r.VacancyID == 0 {
		return errors.New("vacancy_id is required")
	}

	return repository.CreateResponse(r)
}

func GetResponses() ([]models.Response, error) {

	return repository.GetResponses()

}
