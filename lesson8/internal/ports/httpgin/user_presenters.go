package httpgin

import (
	"homework8/internal/models/user"
	"homework8/internal/ports/dto"

	"github.com/gin-gonic/gin"
)

func UserSuccessResponse(usr *user.User) gin.H {
	return gin.H{
		"data": dto.UserResponse{
			ID:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
		},
		"error": nil,
	}
}

func UserErrorResponse(err error) gin.H {
	return gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func ListUsersResponse(userList []user.User) gin.H {
	userResponses := make([]dto.UserResponse, 0, len(userList))
	for _, usr := range userList {
		userResponses = append(userResponses, dto.UserResponse{
			ID:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
		})
	}
	return gin.H{
		"data": userResponses,
	}
}
