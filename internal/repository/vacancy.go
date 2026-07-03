package repository

import (
	"context"

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

	if err != nil {
		return err
	}

	return nil
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