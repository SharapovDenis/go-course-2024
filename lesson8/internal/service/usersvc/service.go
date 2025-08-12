package usersvc

import (
	"errors"
	"homework8/internal/models/enums"
	"homework8/internal/models/user"
	"homework8/internal/repositories/userrepo"
	"regexp"
)

type Service interface {
	GetById(int64) (user.User, error)
	Create(user.User) (user.User, error)
	Update(user.User) (user.User, error)
	List() ([]user.User, error)
}

type service struct {
	repository userrepo.Repository
}

func New(repo userrepo.Repository) Service {
	return &service{repository: repo}
}

func validateUser(usr *user.User) error {
	// Шаблон: <текст>@<текст>.<домен>
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if usr.Name == "" {
		err := errors.New("empty name")
		return errors.Join(enums.ErrValidation, err)
	}

	if usr.Email == "" {
		err := errors.New("empty email")
		return errors.Join(enums.ErrValidation, err)
	}

	if !emailRegex.MatchString(usr.Email) {
		err := errors.New("invalid email format")
		return errors.Join(enums.ErrValidation, err)
	}

	return nil
}

func (s *service) GetById(id int64) (user.User, error) {
	usr, err := s.repository.GetUserById(id)
	if err != nil {
		return user.New(), err
	}
	return usr, nil
}

func (s *service) Create(usr user.User) (user.User, error) {

	err := validateUser(&usr)
	if err != nil {
		return user.New(), err
	}

	usr, err = s.repository.Create(usr)
	if err != nil {
		return user.New(), err
	}

	return usr, nil
}

func (s *service) Update(usr user.User) (user.User, error) {

	err := validateUser(&usr)
	if err != nil {
		return user.New(), err
	}

	usr, err = s.repository.Replace(usr)
	if err != nil {
		return user.New(), err
	}

	return usr, nil

}

func (s *service) List() ([]user.User, error) {
	return s.repository.List()
}
