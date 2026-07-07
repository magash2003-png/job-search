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

	exists, err := repository.VacancyExists(r.VacancyID)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("vacancy not found")
	}

	exists, err = repository.ResponseExists(r.UserID, r.VacancyID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("you have already responded to this vacancy")
	}

	return repository.CreateResponse(r)
}

func GetResponses() ([]models.Response, error) {

	return repository.GetResponses()

}

func GetResponsesByUser(userID int) ([]models.Response, error) {

	return repository.GetResponsesByUser(userID)

}

func GetResponsesByEmployer(employerID int) ([]models.Response, error) {

	return repository.GetResponsesByEmployer(employerID)

}
