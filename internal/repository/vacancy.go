package repository

import (
	"context"
	"errors"

	"job-search/internal/database"
	"job-search/internal/models"
)

func CreateVacancy(v models.Vacancy) error {

	_, err := database.DB.Exec(
		context.Background(),

		`INSERT INTO vacancies
		(title, description, city, salary, employer_id)
		VALUES ($1, $2, $3, $4, $5)`,

		v.Title,
		v.Description,
		v.City,
		v.Salary,
		v.EmployerID,
	)

	return err
}

func GetVacancies() ([]models.Vacancy, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT id, title, description, city, salary, employer_id
		FROM vacancies`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var vacancies []models.Vacancy

	for rows.Next() {

		var vacancy models.Vacancy

		err := rows.Scan(
			&vacancy.ID,
			&vacancy.Title,
			&vacancy.Description,
			&vacancy.City,
			&vacancy.Salary,
			&vacancy.EmployerID,
		)

		if err != nil {
			return nil, err
		}

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func GetVacancyByID(id int) (models.Vacancy, error) {

	var vacancy models.Vacancy

	err := database.DB.QueryRow(
		context.Background(),

		`SELECT id, title, description, city, salary, employer_id
		FROM vacancies
		WHERE id = $1`,

		id,
	).Scan(
		&vacancy.ID,
		&vacancy.Title,
		&vacancy.Description,
		&vacancy.City,
		&vacancy.Salary,
		&vacancy.EmployerID,
	)

	if err != nil {
		return models.Vacancy{}, err
	}

	return vacancy, nil
}

// Employer может изменить только свою вакансию.
func UpdateVacancy(v models.Vacancy) error {

	result, err := database.DB.Exec(
		context.Background(),

		`UPDATE vacancies
		SET title = $1,
		    description = $2,
		    city = $3,
		    salary = $4
		WHERE id = $5
		AND employer_id = $6`,

		v.Title,
		v.Description,
		v.City,
		v.Salary,
		v.ID,
		v.EmployerID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("vacancy not found or access denied")
	}

	return nil
}

// Admin может изменить любую вакансию.
func UpdateVacancyByAdmin(v models.Vacancy) error {

	result, err := database.DB.Exec(
		context.Background(),

		`UPDATE vacancies
		SET title = $1,
		    description = $2,
		    city = $3,
		    salary = $4
		WHERE id = $5`,

		v.Title,
		v.Description,
		v.City,
		v.Salary,
		v.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("vacancy not found")
	}

	return nil
}

// Employer может удалить только свою вакансию.
func DeleteVacancy(id int, employerID int) error {

	result, err := database.DB.Exec(
		context.Background(),

		`DELETE FROM vacancies
		WHERE id = $1
		AND employer_id = $2`,

		id,
		employerID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("vacancy not found or access denied")
	}

	return nil
}

// Admin может удалить любую вакансию.
func DeleteVacancyByAdmin(id int) error {

	result, err := database.DB.Exec(
		context.Background(),

		`DELETE FROM vacancies
		WHERE id = $1`,

		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("vacancy not found")
	}

	return nil
}
