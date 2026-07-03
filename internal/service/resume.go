package service

import (
	"errors"

	"job-search/internal/models"
	"job-search/internal/repository"
)

func CreateResume(r models.Resume) error {

	if r.Title == "" {
		return errors.New("title is required")
	}

	if r.Description == "" {
		return errors.New("description is required")
	}

	return repository.CreateResume(r)
}

func GetResumes() ([]models.Resume, error) {

	return repository.GetResumes()

}