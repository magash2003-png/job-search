package repository

import (
	"context"

	"job-search/internal/database"
	"job-search/internal/models"
)

func CreateFavorite(f models.Favorite) error {

	_, err := database.DB.Exec(
		context.Background(),

		`INSERT INTO favorites (user_id, vacancy_id)
		VALUES ($1, $2)`,

		f.UserID,
		f.VacancyID,
	)

	return err
}

func GetFavorites() ([]models.Favorite, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT id, user_id, vacancy_id
		FROM favorites`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var favorites []models.Favorite

	for rows.Next() {

		var favorite models.Favorite

		err := rows.Scan(
			&favorite.ID,
			&favorite.UserID,
			&favorite.VacancyID,
		)

		if err != nil {
			return nil, err
		}

		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

func DeleteFavorite(id int) error {

	_, err := database.DB.Exec(
		context.Background(),

		`DELETE FROM favorites
		WHERE id = $1`,
		id,
	)

	return err
}