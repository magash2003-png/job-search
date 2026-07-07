package service

import (
	"errors"

	"job-search/internal/models"
	"job-search/internal/repository"
)

func CreateResume(r models.Resume) error {

	if r.FullName == "" {
		return errors.New("full name is required")
	}

	if r.Experience == "" {
		return errors.New("experience is required")
	}

	if r.Skills == "" {
		return errors.New("skills are required")
	}

	if r.Education == "" {
		return errors.New("education is required")
	}

	if r.ExpectedSalary <= 0 {
		return errors.New("expected salary is required")
	}

	return repository.CreateResume(r)
}

func GetResumes() ([]models.Resume, error) {

	return repository.GetResumes()

}

func GetResumesByUser(userID int) ([]models.Resume, error) {

	return repository.GetResumesByUser(userID)

}

func GetResumesForEmployer(employerID int) ([]models.Resume, error) {

	return repository.GetResumesForEmployer(employerID)

}

func UpdateResume(r models.Resume) error {

	if r.ID <= 0 {
		return errors.New("invalid resume id")
	}

	if r.FullName == "" {
		return errors.New("full name is required")
	}

	if r.Experience == "" {
		return errors.New("experience is required")
	}

	if r.Skills == "" {
		return errors.New("skills are required")
	}

	if r.Education == "" {
		return errors.New("education is required")
	}

	if r.ExpectedSalary <= 0 {
		return errors.New("expected salary must be greater than 0")
	}

	return repository.UpdateResume(r)
}

func UpdateResumeByAdmin(r models.Resume) error {

	if r.ID <= 0 {
		return errors.New("invalid resume id")
	}

	if r.FullName == "" {
		return errors.New("full name is required")
	}

	if r.Experience == "" {
		return errors.New("experience is required")
	}

	if r.Skills == "" {
		return errors.New("skills are required")
	}

	if r.Education == "" {
		return errors.New("education is required")
	}

	if r.ExpectedSalary <= 0 {
		return errors.New("expected salary must be greater than 0")
	}

	return repository.UpdateResumeByAdmin(r)
}

func DeleteResume(id int, userID int) error {

	if id <= 0 {
		return errors.New("invalid resume id")
	}

	return repository.DeleteResume(id, userID)
}

func DeleteResumeByAdmin(id int) error {

	if id <= 0 {
		return errors.New("invalid resume id")
	}

	return repository.DeleteResumeByAdmin(id)
}
