package service

import (
	"errors"

	"job-search/internal/models"
	"job-search/internal/repository"
)

func CreateVacancy(v models.Vacancy) error {

	if v.Title == "" {
		return errors.New("title is required")
	}

	if v.Description == "" {
		return errors.New("description is required")
	}

	if v.City == "" {
		return errors.New("city is required")
	}

	if v.Salary <= 0 {
		return errors.New("salary must be greater than 0")
	}

	return repository.CreateVacancy(v)
}

func GetVacancies() ([]models.Vacancy, error) {

	return repository.GetVacancies()

}

func GetVacancyByID(id int) (models.Vacancy, error) {

	return repository.GetVacancyByID(id)

}
