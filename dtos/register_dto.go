package dtos

type RegisterRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
