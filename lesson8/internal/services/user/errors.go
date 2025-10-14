package user

import (
	"errors"
	"fmt"
	repo "homework8/internal/repositories/user"
)

var (
	ErrValidationEmptyName    = fmt.Errorf("empty name: %w", repo.ErrValidation)
	ErrValidationEmptyEmail   = fmt.Errorf("empty email: %w", repo.ErrValidation)
	ErrValidationInvalidEmail = fmt.Errorf("invalid email format: %w", repo.ErrValidation)
	ErrUserNotFound           = fmt.Errorf("user: %w", repo.ErrNotFound)
	ErrUserAlreadyExists      = fmt.Errorf("user: %w", repo.ErrAlreadyExists)
)

func handleRepoError(err error) error {
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return ErrUserNotFound
	case errors.Is(err, repo.ErrAlreadyExists):
		return ErrUserAlreadyExists
	}
	return err
}
