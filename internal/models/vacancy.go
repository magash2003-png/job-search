package models

type Vacancy struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	City        string `json:"city"`
	Salary      int    `json:"salary"`
	EmployerID  int    `json:"employer_id"`
}