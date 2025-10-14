package usecase

import (
	"fmt"
	"homework8/internal/models/user"
	usersvc "homework8/internal/services/user"
)

type CreateUserUseCase struct {
	user usersvc.Service
}

func CreateUser(userSvc usersvc.Service) *CreateUserUseCase {
	return &CreateUserUseCase{user: userSvc}
}

func (u *CreateUserUseCase) Execute(usr user.User) (user.User, error) {
	created, err := u.user.Create(usr)
	if err != nil {
		return user.New(), fmt.Errorf("%w| user=%+v", handleServiceError(err), usr)
	}
	return created, nil
}
