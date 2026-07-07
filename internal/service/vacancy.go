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

func UpdateVacancy(v models.Vacancy) error {

	if v.ID <= 0 {
		return errors.New("invalid vacancy id")
	}

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

	return repository.UpdateVacancy(v)
}

func UpdateVacancyByAdmin(v models.Vacancy) error {

	if v.ID <= 0 {
		return errors.New("invalid vacancy id")
	}

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

	return repository.UpdateVacancyByAdmin(v)
}

func DeleteVacancy(id int, employerID int) error {

	if id <= 0 {
		return errors.New("invalid vacancy id")
	}

	return repository.DeleteVacancy(id, employerID)
}

func DeleteVacancyByAdmin(id int) error {

	if id <= 0 {
		return errors.New("invalid vacancy id")
	}

	return repository.DeleteVacancyByAdmin(id)
}
