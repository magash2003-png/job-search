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

		`SELECT id, user_id, vacancy_id FROM responses`,
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
