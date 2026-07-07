package repository

import (
	"context"

	"job-search/internal/database"
	"job-search/internal/models"
)

func CreateResponse(r models.Response) error {

	_, err := database.DB.Exec(
		context.Background(),

		`INSERT INTO responses (user_id, vacancy_id)
		VALUES ($1, $2)`,

		r.UserID,
		r.VacancyID,
	)

	return err
}

func GetResponses() ([]models.Response, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT id, user_id, vacancy_id
		FROM responses`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var responses []models.Response

	for rows.Next() {

		var response models.Response

		err := rows.Scan(
			&response.ID,
			&response.UserID,
			&response.VacancyID,
		)

		if err != nil {
			return nil, err
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func GetResponsesByUser(userID int) ([]models.Response, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT id, user_id, vacancy_id
		FROM responses
		WHERE user_id = $1`,

		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var responses []models.Response

	for rows.Next() {

		var response models.Response

		err := rows.Scan(
			&response.ID,
			&response.UserID,
			&response.VacancyID,
		)

		if err != nil {
			return nil, err
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func GetResponsesByEmployer(employerID int) ([]models.Response, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT r.id, r.user_id, r.vacancy_id
		FROM responses r
		JOIN vacancies v
		ON r.vacancy_id = v.id
		WHERE v.employer_id = $1`,

		employerID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var responses []models.Response

	for rows.Next() {

		var response models.Response

		err := rows.Scan(
			&response.ID,
			&response.UserID,
			&response.VacancyID,
		)

		if err != nil {
			return nil, err
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func VacancyExists(vacancyID int) (bool, error) {

	var count int

	err := database.DB.QueryRow(
		context.Background(),

		`SELECT COUNT(*)
		FROM vacancies
		WHERE id = $1`,

		vacancyID,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func ResponseExists(userID int, vacancyID int) (bool, error) {

	var count int

	err := database.DB.QueryRow(
		context.Background(),

		`SELECT COUNT(*)
		FROM responses
		WHERE user_id = $1
		AND vacancy_id = $2`,

		userID,
		vacancyID,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
