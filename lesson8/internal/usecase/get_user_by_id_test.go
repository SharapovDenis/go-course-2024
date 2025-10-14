package usecase

import (
	"homework8/internal/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_GetUserById(t *testing.T) {
	mockUserSvc := new(MockUserService)
	uc := GetUserById(mockUserSvc)

	usr := user.User{ID: 1, Name: "John"}
	mockUserSvc.On("GetById", int64(1)).Return(usr, nil)

	got, err := uc.Execute(1)
	assert.NoError(t, err)
	assert.Equal(t, usr, got)

	mockUserSvc.AssertExpectations(t)
}
