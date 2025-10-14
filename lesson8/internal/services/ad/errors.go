package ad

import (
	"errors"
	"fmt"
	repo "homework8/internal/repositories/ad"
)

var (
	ErrValidationEmptyTitle   = fmt.Errorf("empty title: %w", repo.ErrValidation)
	ErrValidationTooLongTitle = fmt.Errorf("title is too long: %w", repo.ErrValidation)
	ErrValidationEmptyText    = fmt.Errorf("empty text: %w", repo.ErrValidation)
	ErrValidationTooLongText  = fmt.Errorf("text is too long: %w", repo.ErrValidation)
	ErrAdNotFound             = fmt.Errorf("ad: %w", repo.ErrNotFound)
	ErrAdAlreadyExists        = fmt.Errorf("ad: %w", repo.ErrAlreadyExists)
)

func handleRepoError(err error) error {
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return ErrAdNotFound
	case errors.Is(err, repo.ErrAlreadyExists):
		return ErrAdAlreadyExists
	}
	return err
}
