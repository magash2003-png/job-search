package models

type Resume struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	FullName       string `json:"full_name"`
	Experience     string `json:"experience"`
	Skills         string `json:"skills"`
	Education      string `json:"education"`
	ExpectedSalary int    `json:"expected_salary"`
}
