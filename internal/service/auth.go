package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"job-search/internal/auth"
	"job-search/internal/models"
	"job-search/internal/repository"
)

func RegisterUser(user models.User) error {

	if user.Username == "" {
		return errors.New("username is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if user.Role == "admin" {
		return errors.New("admin role cannot be assigned during registration")
	}

	if user.Role != "applicant" && user.Role != "employer" {
		return errors.New("role must be applicant or employer")
	}

	_, err := repository.FindUserByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return repository.CreateUser(user)
}

func LoginUser(email string, password string) (string, error) {

	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ChangeUserRole(userID int, role string) error {

	if role != "admin" &&
		role != "employer" &&
		role != "applicant" {
		return errors.New("invalid role")
	}

	return repository.ChangeUserRole(userID, role)
}
