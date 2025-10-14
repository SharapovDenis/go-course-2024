package user

import (
	"homework8/internal/models/user"
)

type Repository interface {
	GetUserById(int64) (user.User, error)
	Create(user.User) (user.User, error)
	Replace(user.User) (user.User, error)
	List() ([]user.User, error)
}
