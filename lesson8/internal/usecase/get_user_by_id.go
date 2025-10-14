package usecase

import (
	"fmt"
	"homework8/internal/models/user"
	usersvc "homework8/internal/services/user"
)

type GetUserByIdUseCase struct {
	user usersvc.Service
}

func GetUserById(userSvc usersvc.Service) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{user: userSvc}
}

func (u *GetUserByIdUseCase) Execute(id int64) (user.User, error) {
	found, err := u.user.GetById(id)
	if err != nil {
		return user.New(), fmt.Errorf("%w| id=%d", handleServiceError(err), id)
	}
	return found, nil
}
