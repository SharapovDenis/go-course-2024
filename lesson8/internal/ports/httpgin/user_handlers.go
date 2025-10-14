package httpgin

import (
	"errors"
	"homework8/internal/ports/dto"
	usersvc "homework8/internal/services/user"
	"homework8/internal/usecase"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func handleUserError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, usecase.Err4003_001):
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
	case errors.Is(err, usecase.Err4003_002):
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
	case errors.Is(err, usecase.Err4003_003):
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
	default:
		c.JSON(http.StatusInternalServerError, UserErrorResponse(err))
	}
}

func createUser(userSvc usersvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody dto.CreateUserRequest
		if err := c.BindJSON(&reqBody); err != nil {
			handleUserError(c, err)
			return
		}
		uc := usecase.CreateUser(userSvc)
		usr, err := uc.Execute(reqBody.ToUser())
		if err != nil {
			handleUserError(c, err)
			return
		}
		c.JSON(http.StatusOK, UserSuccessResponse(&usr))
	}
}

func updateUser(userSvc usersvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody dto.UpdateUserRequest
		if err := c.BindJSON(&reqBody); err != nil {
			handleUserError(c, err)
			return
		}

		strUserID := c.Param("user_id")
		if strUserID == "" {
			handleUserError(c, errors.New("empty user_id"))
			return
		}

		intUserID, err := strconv.ParseInt(strUserID, 10, 64)
		if err != nil {
			handleUserError(c, err)
			return
		}

		uc := usecase.UpdateUser(userSvc)
		usr, err := uc.Execute(intUserID, reqBody.Name, reqBody.Email)
		if err != nil {
			handleUserError(c, err)
			return
		}
		c.JSON(http.StatusOK, UserSuccessResponse(&usr))
	}
}

func listUsers(userSvc usersvc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		uc := usecase.ListUsers(userSvc)
		userList, err := uc.Execute()
		if err != nil {
			handleUserError(c, err)
			return
		}
		c.JSON(http.StatusOK, ListUsersResponse(userList))
	}
}
