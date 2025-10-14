package usecase

import (
	"homework8/internal/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_UpdateUser(t *testing.T) {
	mockUserSvc := new(MockUserService)
	uc := UpdateUser(mockUserSvc)

	usr := user.User{ID: 1, Name: "John", Email: "john@example.com"}
	updated := user.User{ID: 1, Name: "John Updated", Email: "john.updated@example.com"}

	mockUserSvc.On("GetById", usr.ID).Return(usr, nil)

	mockUserSvc.On("Update", updated).Return(updated, nil)

	got, err := uc.Execute(updated.ID, updated.Name, updated.Email)

	assert.NoError(t, err)
	assert.Equal(t, updated, got)

	// Проверяем, что все ожидания выполнились
	mockUserSvc.AssertExpectations(t)
}
