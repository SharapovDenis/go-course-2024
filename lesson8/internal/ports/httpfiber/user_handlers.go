package httpfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"errors"
	"homework8/internal/models/enums"
	"homework8/internal/ports/dto"
	"homework8/internal/service"
)

func handleUserError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, enums.ErrValidation):
		c.Status(http.StatusBadRequest)
		return c.JSON(UserErrorResponse(err))
	case errors.Is(err, enums.ErrForbidden):
		c.Status(http.StatusForbidden)
		return c.JSON(UserErrorResponse(err))
	default:
		c.Status(http.StatusInternalServerError)
		return c.JSON(UserErrorResponse(err))
	}
}

func createUser(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody dto.CreateUserRequest
		err := c.BodyParser(&reqBody)
		if err != nil {
			return handleUserError(c, err)
		}
		usr, err := svc.CreateUser(reqBody.ToUser())
		if err != nil {
			return handleUserError(c, err)
		}
		return c.JSON(UserSuccessResponse(&usr))
	}
}

func updateUser(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody dto.UpdateUserRequest
		if err := c.BodyParser(&reqBody); err != nil {
			return handleUserError(c, err)
		}

		usrID, err := c.ParamsInt("user_id")
		if err != nil {
			return handleUserError(c, err)
		}

		usr, err := svc.UpdateUser(int64(usrID), reqBody.Name, reqBody.Email)
		if err != nil {
			return handleUserError(c, err)
		}

		return c.JSON(UserSuccessResponse(&usr))
	}
}

func listUsers(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		usrList, err := svc.ListUsers()
		if err != nil {
			return handleUserError(c, err)
		}
		return c.JSON(ListUserResponse(usrList))
	}
}
