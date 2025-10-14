package user

import (
	"homework8/internal/models/user"
	repo "homework8/internal/repositories/user"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мок репозитория
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetUserById(id int64) (user.User, error) {
	args := m.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockRepo) Create(usr user.User) (user.User, error) {
	args := m.Called(usr)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockRepo) Replace(usr user.User) (user.User, error) {
	args := m.Called(usr)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockRepo) List() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}

func TestPositive_Create(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	usr := user.User{Name: "John Doe", Email: "john@example.com"}

	mockRepo.On("Create", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(user.User)
		arg.ID = 1
		args[0] = arg
	}).Return(user.User{ID: 1, Name: usr.Name, Email: usr.Email}, nil)

	created, err := svc.Create(usr)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), created.ID)
	assert.Equal(t, usr.Name, created.Name)
	assert.Equal(t, usr.Email, created.Email)

	mockRepo.AssertExpectations(t)
}

func TestNegative_Create_ValidationFails_EmptyName(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	cases := []user.User{
		{Name: "", Email: "john@example.com"},
	}

	for _, c := range cases {
		created, err := svc.Create(c)
		assert.ErrorIs(t, err, ErrValidationEmptyName)
		assert.Equal(t, user.User{}, created)
	}
}

func TestNegative_Create_ValidationFails_EmptyEmail(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	cases := []user.User{
		{Name: "John", Email: ""},
	}

	for _, c := range cases {
		created, err := svc.Create(c)
		assert.ErrorIs(t, err, ErrValidationEmptyEmail)
		assert.Equal(t, user.User{}, created)
	}
}

func TestNegative_Create_ValidationFails_InvalidEmail(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	cases := []user.User{
		{Name: "John", Email: "invalid-email"},
	}

	for _, c := range cases {
		created, err := svc.Create(c)
		assert.ErrorIs(t, err, ErrValidationInvalidEmail)
		assert.Equal(t, user.User{}, created)
	}
}

func TestPositive_Update(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	usr := user.User{ID: 1, Name: "John Updated", Email: "john.updated@example.com"}

	mockRepo.On("Replace", usr).Return(usr, nil)

	updated, err := svc.Update(usr)
	assert.NoError(t, err)
	assert.Equal(t, usr, updated)

	mockRepo.AssertExpectations(t)
}

func TestNegative_Update_ValidationFails_EmptyName(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	cases := []user.User{
		{ID: 1, Name: "", Email: "john@example.com"},
	}

	for _, c := range cases {
		updated, err := svc.Update(c)
		assert.ErrorIs(t, err, ErrValidationEmptyName)
		assert.Equal(t, user.User{}, updated)
	}
}

func TestNegative_Update_ValidationFails_EmptyEmail(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	cases := []user.User{
		{ID: 1, Name: "John", Email: ""},
	}

	for _, c := range cases {
		updated, err := svc.Update(c)
		assert.ErrorIs(t, err, ErrValidationEmptyEmail)
		assert.Equal(t, user.User{}, updated)
	}
}

func TestNegative_Update_ValidationFails_InvalidEmail(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	cases := []user.User{
		{ID: 1, Name: "John", Email: "invalid-email"},
	}

	for _, c := range cases {
		updated, err := svc.Update(c)
		assert.ErrorIs(t, err, ErrValidationInvalidEmail)
		assert.Equal(t, user.User{}, updated)
	}
}

func TestPositive_GetById(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	usr := user.User{ID: 1, Name: "John", Email: "john@example.com"}

	mockRepo.On("GetUserById", int64(1)).Return(usr, nil)

	got, err := svc.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, usr, got)

	mockRepo.AssertExpectations(t)
}

func TestNegative_GetById_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	mockRepo.On("GetUserById", int64(1)).Return(user.User{}, repo.ErrNotFound)

	got, err := svc.GetById(1)
	assert.ErrorIs(t, err, ErrUserNotFound)
	assert.Equal(t, user.User{}, got)

	mockRepo.AssertExpectations(t)
}

func TestPositive_List(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	list := []user.User{
		{ID: 1, Name: "John", Email: "john@example.com"},
		{ID: 2, Name: "Jane", Email: "jane@example.com"},
	}

	mockRepo.On("List").Return(list, nil)

	got, err := svc.List()
	assert.NoError(t, err)
	assert.Equal(t, list, got)

	mockRepo.AssertExpectations(t)
}
