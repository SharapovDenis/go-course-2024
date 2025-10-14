package usecase

import (
	"homework8/internal/models/user"
	usersvc "homework8/internal/services/user"
)

type ListUsersUseCase struct {
	user usersvc.Service
}

func ListUsers(userSvc usersvc.Service) *ListUsersUseCase {
	return &ListUsersUseCase{user: userSvc}
}

func (u *ListUsersUseCase) Execute() ([]user.User, error) {
	return u.user.List()
}
