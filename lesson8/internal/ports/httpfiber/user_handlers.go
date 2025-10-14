package httpfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"errors"
	"homework8/internal/ports/dto"
	usersvc "homework8/internal/services/user"
	"homework8/internal/usecase"
)

func handleUserError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, usecase.Err4003_001):
		c.Status(http.StatusBadRequest)
		return c.JSON(UserErrorResponse(err))
	case errors.Is(err, usecase.Err4003_002):
		c.Status(http.StatusBadRequest)
		return c.JSON(UserErrorResponse(err))
	case errors.Is(err, usecase.Err4003_003):
		c.Status(http.StatusBadRequest)
		return c.JSON(UserErrorResponse(err))
	case errors.Is(err, usecase.Err4001_002):
		c.Status(http.StatusForbidden)
		return c.JSON(UserErrorResponse(err))
	default:
		c.Status(http.StatusInternalServerError)
		return c.JSON(UserErrorResponse(err))
	}
}

func createUser(userSvc usersvc.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody dto.CreateUserRequest
		err := c.BodyParser(&reqBody)
		if err != nil {
			return handleUserError(c, err)
		}
		uc := usecase.CreateUser(userSvc)
		usr, err := uc.Execute(reqBody.ToUser())
		if err != nil {
			return handleUserError(c, err)
		}
		return c.JSON(UserSuccessResponse(&usr))
	}
}

func updateUser(userSvc usersvc.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody dto.UpdateUserRequest
		if err := c.BodyParser(&reqBody); err != nil {
			return handleUserError(c, err)
		}

		usrID, err := c.ParamsInt("user_id")
		if err != nil {
			return handleUserError(c, err)
		}

		uc := usecase.UpdateUser(userSvc)
		usr, err := uc.Execute(int64(usrID), reqBody.Name, reqBody.Email)
		if err != nil {
			return handleUserError(c, err)
		}

		return c.JSON(UserSuccessResponse(&usr))
	}
}

func listUsers(userSvc usersvc.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		uc := usecase.ListUsers(userSvc)
		usrList, err := uc.Execute()
		if err != nil {
			return handleUserError(c, err)
		}
		return c.JSON(ListUserResponse(usrList))
	}
}
