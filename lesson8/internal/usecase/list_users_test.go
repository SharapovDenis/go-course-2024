package usecase

import (
	"homework8/internal/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_ListUsers(t *testing.T) {
	mockUserSvc := new(MockUserService)
	uc := ListUsers(mockUserSvc)

	users := []user.User{{ID: 1}, {ID: 2}}
	mockUserSvc.On("List").Return(users, nil)

	got, err := uc.Execute()
	assert.NoError(t, err)
	assert.Equal(t, users, got)

	mockUserSvc.AssertExpectations(t)
}
