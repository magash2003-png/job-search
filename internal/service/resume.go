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
