package service_test

import (
	"homework8/internal/models/ads"
	"homework8/internal/models/enums"
	"homework8/internal/models/user"
	"homework8/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPositive_GetById(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	ad := ads.Ad{ID: 1, Title: "title", AuthorID: 2}
	mockAdSvc.On("GetById", int64(1)).Return(ad, nil)

	got, err := svc.GetAdById(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, ad, got)

	mockAdSvc.AssertExpectations(t)
}

func TestPositive_CreateAd(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	usr := user.User{ID: 2, Name: "Alice"}

	ad := ads.Ad{Title: "title", Text: "text", AuthorID: usr.ID}

	created := ad
	created.ID = 1

	mockUserSvc.On("GetById", usr.ID).Return(usr, nil)
	mockAdSvc.On("Create", ad).Return(created, nil)

	got, err := svc.CreateAd(ad, usr.ID)

	assert.NoError(t, err)
	assert.Equal(t, created, got)

	mockUserSvc.AssertExpectations(t)
	mockAdSvc.AssertExpectations(t)
}

func TestPositive_UpdateAd(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	origAd := ads.Ad{ID: 1, Title: "old", Text: "old text", AuthorID: 2}
	updatedAd := origAd
	updatedAd.Title = "new"
	updatedAd.Text = "new text"

	mockAdSvc.On("GetById", int64(1)).Return(origAd, nil)
	mockAdSvc.On("Update", mock.MatchedBy(func(ad ads.Ad) bool {
		return ad.ID == 1 && ad.Title == "new" && ad.Text == "new text"
	})).Return(updatedAd, nil)

	got, err := svc.UpdateAd(1, "new", "new text", 2)
	assert.NoError(t, err)
	assert.Equal(t, updatedAd, got)

	mockAdSvc.AssertExpectations(t)
}

func TestNegative_UpdateAd_Forbidden(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	origAd := ads.Ad{ID: 1, Title: "old", AuthorID: 2}

	mockAdSvc.On("GetById", int64(1)).Return(origAd, nil)

	_, err := svc.UpdateAd(1, "new", "new text", 3) // userID != authorID
	assert.ErrorIs(t, err, enums.ErrForbidden)

	mockAdSvc.AssertExpectations(t)
}

func TestNegative_UpdateAd_NotFound(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	mockAdSvc.On("GetById", int64(1)).Return(ads.Ad{}, enums.ErrNotFound)

	_, err := svc.UpdateAd(1, "new", "new text", 2)
	assert.ErrorIs(t, err, enums.ErrNotFound)

	mockAdSvc.AssertExpectations(t)
}

func TestPositive_ChangeAdStatus(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	ad := ads.Ad{ID: 1, Published: false, AuthorID: 2}
	updatedAd := ad
	updatedAd.Published = true

	mockAdSvc.On("GetById", int64(1)).Return(ad, nil)
	mockAdSvc.On("Update", mock.MatchedBy(func(ad ads.Ad) bool {
		return ad.ID == 1 && ad.Published == true
	})).Return(updatedAd, nil)

	got, err := svc.ChangeAdStatus(1, true, 2)
	assert.NoError(t, err)
	assert.Equal(t, updatedAd, got)

	mockAdSvc.AssertExpectations(t)
}

func TestNegative_ChangeAdStatus_Forbidden(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	ad := ads.Ad{ID: 1, Published: false, AuthorID: 2}

	mockAdSvc.On("GetById", int64(1)).Return(ad, nil)

	_, err := svc.ChangeAdStatus(1, true, 3) // userID != authorID
	assert.ErrorIs(t, err, enums.ErrForbidden)

	mockAdSvc.AssertExpectations(t)
}

func TestNegative_ChangeAdStatus_NotFound(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	mockAdSvc.On("GetById", int64(1)).Return(ads.Ad{}, enums.ErrNotFound)

	_, err := svc.ChangeAdStatus(1, true, 2)
	assert.ErrorIs(t, err, enums.ErrNotFound)

	mockAdSvc.AssertExpectations(t)
}

func TestPositive_GetUserById(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	usr := user.User{ID: 1, Name: "John"}
	mockUserSvc.On("GetById", int64(1)).Return(usr, nil)

	got, err := svc.GetUserById(1)
	assert.NoError(t, err)
	assert.Equal(t, usr, got)

	mockUserSvc.AssertExpectations(t)
}

func TestPositive_CreateUser(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	usr := user.User{Name: "John"}
	created := usr
	created.ID = 1

	mockUserSvc.On("Create", usr).Return(created, nil)

	got, err := svc.CreateUser(usr)
	assert.NoError(t, err)
	assert.Equal(t, created, got)

	mockUserSvc.AssertExpectations(t)
}

func TestPositive_UpdateUser(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	usr := user.User{ID: 1, Name: "John", Email: "john@example.com"}
	updated := user.User{ID: 1, Name: "John Updated", Email: "john.updated@example.com"}

	mockUserSvc.On("GetById", usr.ID).Return(usr, nil)

	mockUserSvc.On("Update", updated).Return(updated, nil)

	got, err := svc.UpdateUser(updated.ID, updated.Name, updated.Email)

	assert.NoError(t, err)
	assert.Equal(t, updated, got)

	// Проверяем, что все ожидания выполнились
	mockUserSvc.AssertExpectations(t)
}

func TestPositive_ListUsers(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	users := []user.User{{ID: 1}, {ID: 2}}
	mockUserSvc.On("List").Return(users, nil)

	got, err := svc.ListUsers()
	assert.NoError(t, err)
	assert.Equal(t, users, got)

	mockUserSvc.AssertExpectations(t)
}

func TestPositive_ListAd(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	svc := service.New(mockAdSvc, mockUserSvc)

	adsList := []ads.Ad{{ID: 1}, {ID: 2}}
	filter := ads.NewFilter()
	mockAdSvc.On("List", filter).Return(adsList, nil)

	got, err := svc.ListAd(filter)
	assert.NoError(t, err)
	assert.Equal(t, adsList, got)

	mockAdSvc.AssertExpectations(t)
}
