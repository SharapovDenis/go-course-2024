package dto

import "homework8/internal/models/user"

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *CreateUserRequest) ToUser() user.User {
	usr := user.New()
	usr.Name = r.Name
	usr.Email = r.Email
	return usr
}
