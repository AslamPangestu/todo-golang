package dtos

import "todo-be/entities"

type AuthResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

func AuthAdapter(user entities.User, token string) AuthResponse {
	return AuthResponse{
		ID:     user.ID,
		Name:   user.Name,
		UserID: user.UserID,
		Token:  token,
	}
}
