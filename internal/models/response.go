package models

type Response struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	VacancyID  int `json:"vacancy_id"`
}