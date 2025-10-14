package user

import (
	"homework8/internal/models/user"
	repo "homework8/internal/repositories/user"
	"regexp"
)

type Service interface {
	GetById(int64) (user.User, error)
	Create(user.User) (user.User, error)
	Update(user.User) (user.User, error)
	List() ([]user.User, error)
}

type service struct {
	repository repo.Repository
}

func New(repo repo.Repository) Service {
	return &service{repository: repo}
}

func validateUser(usr *user.User) error {
	// Шаблон: <текст>@<текст>.<домен>
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if usr.Name == "" {
		return ErrValidationEmptyName
	}

	if usr.Email == "" {
		return ErrValidationEmptyEmail
	}

	if !emailRegex.MatchString(usr.Email) {
		return ErrValidationInvalidEmail
	}

	return nil
}

func (s *service) GetById(id int64) (user.User, error) {
	usr, err := s.repository.GetUserById(id)
	if err != nil {
		return user.New(), handleRepoError(err)
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
		return user.New(), handleRepoError(err)
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
		return user.New(), handleRepoError(err)
	}

	return usr, nil

}

func (s *service) List() ([]user.User, error) {
	return s.repository.List()
}
