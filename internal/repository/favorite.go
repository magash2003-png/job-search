package repository

import (
	"context"
	"errors"

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

func GetFavorites(userID int) ([]models.Favorite, error) {

	rows, err := database.DB.Query(
		context.Background(),

		`SELECT id, user_id, vacancy_id
		FROM favorites
		WHERE user_id = $1`,

		userID,
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

func DeleteFavorite(id int, userID int) error {

	result, err := database.DB.Exec(
		context.Background(),

		`DELETE FROM favorites
		WHERE id = $1
		AND user_id = $2`,

		id,
		userID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("favorite not found")
	}

	return nil
}


func FavoriteExists(userID int, vacancyID int) (bool, error) {

	var count int

	err := database.DB.QueryRow(
		context.Background(),

		`SELECT COUNT(*)
		FROM favorites
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
