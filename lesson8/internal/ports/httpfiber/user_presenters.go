// Здесь делаем адаптацию фреймворка к общему dto

package httpfiber

import (
	"github.com/gofiber/fiber/v2"

	"homework8/internal/models/user"
	"homework8/internal/ports/dto"
)

func UserSuccessResponse(usr *user.User) *fiber.Map {
	return &fiber.Map{
		"data": dto.UserResponse{
			ID:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
		},
		"error": nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  nil,
		"error": err.Error(),
	}
}

func ListUserResponse(usrList []user.User) *fiber.Map {
	usrResponses := make([]dto.UserResponse, 0, len(usrList))
	for _, usr := range usrList {
		usrResponses = append(usrResponses, dto.UserResponse{
			ID:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
		})
	}
	return &fiber.Map{
		"data": usrResponses,
	}
}
