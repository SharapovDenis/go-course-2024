package usecase

import (
	"errors"
	"fmt"
	adsvc "homework8/internal/services/ad"
	usersvc "homework8/internal/services/user"
)

var (
	Err4001_001 = errors.New("4001-001| Произошла непредвиденная ошибка")
	Err4001_002 = errors.New("4001-002| Изменения доступны только автору")

	Err4002_001 = errors.New("4002-001| Публикация не найдена")
	Err4002_002 = errors.New("4002-002| Публикация уже существует")
	Err4002_003 = errors.New("4002-003| Ошибка валидации публикации")

	Err4003_001 = errors.New("4002-001| Пользователь не найден")
	Err4003_002 = errors.New("4002-002| Пользователь уже существует")
	Err4003_003 = errors.New("4002-003| Ошибка валидации пользователя")
)

func handleServiceError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, adsvc.ErrAdNotFound):
		return Err4002_001
	case errors.Is(err, adsvc.ErrAdAlreadyExists):
		return Err4002_002
	case errors.Is(err, adsvc.ErrValidationEmptyTitle):
		return Err4002_003
	case errors.Is(err, adsvc.ErrValidationTooLongTitle):
		return Err4002_003
	case errors.Is(err, adsvc.ErrValidationEmptyText):
		return Err4002_003
	case errors.Is(err, adsvc.ErrValidationTooLongText):
		return Err4002_003
	case errors.Is(err, usersvc.ErrUserNotFound):
		return Err4003_001
	case errors.Is(err, usersvc.ErrUserAlreadyExists):
		return Err4003_002
	case errors.Is(err, usersvc.ErrValidationEmptyName):
		return Err4003_003
	case errors.Is(err, usersvc.ErrValidationEmptyEmail):
		return Err4003_003
	case errors.Is(err, usersvc.ErrValidationInvalidEmail):
		return Err4003_003
	}
	return fmt.Errorf("%w| ", err)
}
