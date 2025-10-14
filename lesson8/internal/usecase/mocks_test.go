package usecase

import (
	"homework8/internal/models/ads"
	"homework8/internal/models/user"

	"github.com/stretchr/testify/mock"
)

type MockAdService struct {
	mock.Mock
}

func (m *MockAdService) GetById(id int64) (ads.Ad, error) {
	args := m.Called(id)
	return args.Get(0).(ads.Ad), args.Error(1)
}

func (m *MockAdService) Create(ad ads.Ad) (ads.Ad, error) {
	args := m.Called(ad)
	return args.Get(0).(ads.Ad), args.Error(1)
}

func (m *MockAdService) Update(ad ads.Ad) (ads.Ad, error) {
	args := m.Called(ad)
	return args.Get(0).(ads.Ad), args.Error(1)
}

func (m *MockAdService) List(f *ads.Filter) ([]ads.Ad, error) {
	args := m.Called(f)
	return args.Get(0).([]ads.Ad), args.Error(1)
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetById(id int64) (user.User, error) {
	args := m.Called(id)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockUserService) Create(usr user.User) (user.User, error) {
	args := m.Called(usr)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockUserService) Update(usr user.User) (user.User, error) {
	args := m.Called(usr)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockUserService) List() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}
