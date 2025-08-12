package httpgin

import (
	"errors"
	"homework8/internal/models/enums"
	"homework8/internal/ports/dto"
	"homework8/internal/service"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func handleUserError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, enums.ErrValidation):
		c.JSON(http.StatusBadRequest, UserErrorResponse(err))
	default:
		c.JSON(http.StatusInternalServerError, UserErrorResponse(err))
	}
}

func createUser(svc service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody dto.CreateUserRequest
		if err := c.BindJSON(&reqBody); err != nil {
			handleUserError(c, err)
			return
		}
		usr, err := svc.CreateUser(reqBody.ToUser())
		if err != nil {
			handleUserError(c, err)
			return
		}
		c.JSON(http.StatusOK, UserSuccessResponse(&usr))
	}
}

func updateUser(svc service.Service) gin.HandlerFunc {
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

		usr, err := svc.UpdateUser(intUserID, reqBody.Name, reqBody.Email)
		if err != nil {
			handleUserError(c, err)
			return
		}
		c.JSON(http.StatusOK, UserSuccessResponse(&usr))
	}
}

func listUsers(svc service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userList, err := svc.ListUsers()
		if err != nil {
			handleUserError(c, err)
			return
		}
		c.JSON(http.StatusOK, ListUsersResponse(userList))
	}
}
