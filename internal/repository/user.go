package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"job-search/internal/database"
	"job-search/internal/models"
)

func CreateUser(user models.User) error {

	_, err := database.DB.Exec(
		context.Background(),

		`INSERT INTO users (username, email, password, role)
		 VALUES ($1, $2, $3, $4)`,

		user.Username,
		user.Email,
		user.Password,
		user.Role,
	)

	if err != nil {

		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return errors.New("email already exists")
			}
		}

		return err
	}

	return nil
}

func FindUserByEmail(email string) (models.User, error) {

	var user models.User

	err := database.DB.QueryRow(
		context.Background(),

		`SELECT id, username, email, password, role
		 FROM users
		 WHERE email = $1`,

		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func ChangeUserRole(userID int, role string) error {

	tag, err := database.DB.Exec(
		context.Background(),

		`UPDATE users
		SET role = $1
		WHERE id = $2`,

		role,
		userID,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}
