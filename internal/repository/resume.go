package repository

import (
	"context"

	"job-search/internal/database"
	"job-search/internal/models"
)

func CreateResume(r models.Resume) error {

	_, err := database.DB.Exec(
		context.Background(),

		`INSERT INTO resumes
		(title, description, skills, user_id)
		VALUES ($1, $2, $3, $4)`,

		r.Title,
		r.Description,
		r.Skills,
		r.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetResumes() ([]models.Resume, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT id, title, description, skills, user_id
		FROM resumes`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resumes []models.Resume

	for rows.Next() {

		var resume models.Resume

		err := rows.Scan(
			&resume.ID,
			&resume.Title,
			&resume.Description,
			&resume.Skills,
			&resume.UserID,
		)

		if err != nil {
			return nil, err
		}

		resumes = append(resumes, resume)
	}

	return resumes, nil
}