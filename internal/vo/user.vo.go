package vo

import "time"

type CreateUserRequest struct {
	UserName string `json:"user_name" binding:"required,alphanum,min=6"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type CreateUserResponse struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt time.Time `json:"created_at"`
}