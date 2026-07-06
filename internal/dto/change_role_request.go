package dto

type ChangeRoleRequest struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}
