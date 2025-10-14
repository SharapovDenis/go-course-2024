package usecase

import (
	"homework8/internal/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_CreateUser(t *testing.T) {
	mockUserSvc := new(MockUserService)
	uc := CreateUser(mockUserSvc)

	usr := user.User{Name: "John"}
	created := usr
	created.ID = 1

	mockUserSvc.On("Create", usr).Return(created, nil)

	got, err := uc.Execute(usr)
	assert.NoError(t, err)
	assert.Equal(t, created, got)

	mockUserSvc.AssertExpectations(t)
}
