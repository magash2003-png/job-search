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
		(full_name, experience, skills, education, expected_salary, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)`,

		r.FullName,
		r.Experience,
		r.Skills,
		r.Education,
		r.ExpectedSalary,
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

		`SELECT id, user_id, full_name, experience, skills, education, expected_salary
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
			&resume.UserID,
			&resume.FullName,
			&resume.Experience,
			&resume.Skills,
			&resume.Education,
			&resume.ExpectedSalary,
		)

		if err != nil {
			return nil, err
		}

		resumes = append(resumes, resume)
	}

	return resumes, nil
}
