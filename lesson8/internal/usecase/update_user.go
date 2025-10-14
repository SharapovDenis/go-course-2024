package usecase

import (
	"fmt"
	"homework8/internal/models/user"
	usersvc "homework8/internal/services/user"
)

type UpdateUserUseCase struct {
	user usersvc.Service
}

func UpdateUser(userSvc usersvc.Service) *UpdateUserUseCase {
	return &UpdateUserUseCase{user: userSvc}
}

func (u *UpdateUserUseCase) Execute(userID int64, name string, email string) (user.User, error) {

	// Делаем поиск пользователя
	usr, err := u.user.GetById(userID)
	if err != nil {
		return user.New(), err
	}

	usr.Name = name
	usr.Email = email

	updated, err := u.user.Update(usr)
	if err != nil {
		return user.New(), fmt.Errorf("%w| userID=%d, name=%s, email=%s", handleServiceError(err), userID, name, email)
	}

	return updated, nil
}
