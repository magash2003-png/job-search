package repository

import (
	"context"
	"errors"

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

func GetResumesByUser(userID int) ([]models.Resume, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT id, user_id, full_name, experience, skills, education, expected_salary
		FROM resumes
		WHERE user_id = $1`,

		userID,
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

func UpdateResume(r models.Resume) error {

	result, err := database.DB.Exec(
		context.Background(),

		`UPDATE resumes
		SET full_name = $1,
		    experience = $2,
		    skills = $3,
		    education = $4,
		    expected_salary = $5
		WHERE id = $6
		AND user_id = $7`,

		r.FullName,
		r.Experience,
		r.Skills,
		r.Education,
		r.ExpectedSalary,
		r.ID,
		r.UserID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("resume not found or access denied")
	}

	return nil
}

func UpdateResumeByAdmin(r models.Resume) error {

	result, err := database.DB.Exec(
		context.Background(),

		`UPDATE resumes
		SET full_name = $1,
		    experience = $2,
		    skills = $3,
		    education = $4,
		    expected_salary = $5
		WHERE id = $6`,

		r.FullName,
		r.Experience,
		r.Skills,
		r.Education,
		r.ExpectedSalary,
		r.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("resume not found")
	}

	return nil
}

func DeleteResume(id int, userID int) error {

	result, err := database.DB.Exec(
		context.Background(),

		`DELETE FROM resumes
		WHERE id = $1
		AND user_id = $2`,

		id,
		userID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("resume not found or access denied")
	}

	return nil
}

func DeleteResumeByAdmin(id int) error {

	result, err := database.DB.Exec(
		context.Background(),

		`DELETE FROM resumes
		WHERE id = $1`,

		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("resume not found")
	}

	return nil
}

func GetResumesForEmployer(employerID int) ([]models.Resume, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT DISTINCT
			r.id,
			r.user_id,
			r.full_name,
			r.experience,
			r.skills,
			r.education,
			r.expected_salary
		FROM resumes r
		INNER JOIN responses resp
			ON r.user_id = resp.user_id
		INNER JOIN vacancies v
			ON resp.vacancy_id = v.id
		WHERE v.employer_id = $1`,

		employerID,
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
