package ad

import (
	"homework8/internal/models/ads"
	adrepo "homework8/internal/repositories/ad"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мок репозитория
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetAdById(id int64) (ads.Ad, error) {
	args := m.Called(id)
	return args.Get(0).(ads.Ad), args.Error(1)
}

func (m *MockRepo) Create(ad ads.Ad) (ads.Ad, error) {
	args := m.Called(ad)
	return args.Get(0).(ads.Ad), args.Error(1)
}

func (m *MockRepo) Replace(ad ads.Ad) (ads.Ad, error) {
	args := m.Called(ad)
	return args.Get(0).(ads.Ad), args.Error(1)
}

func (m *MockRepo) List(f *ads.Filter) ([]ads.Ad, error) {
	args := m.Called(f)
	return args.Get(0).([]ads.Ad), args.Error(1)
}

func TestNegative_Create_ValidationFails(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	ad := ads.Ad{Title: "", Text: "text"}
	_, err := svc.Create(ad)
	assert.ErrorIs(t, err, ErrValidationEmptyTitle)

	ad = ads.Ad{Title: "valid title", Text: ""}
	_, err = svc.Create(ad)
	assert.ErrorIs(t, err, ErrValidationEmptyText)
}

func TestPositive_Create_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	ad := ads.Ad{Title: "valid title", Text: "valid text"}

	mockRepo.On("Create", mock.MatchedBy(func(a ads.Ad) bool {
		return a.Title == ad.Title && a.Text == ad.Text
	})).Return(ads.Ad{ID: 1, Title: ad.Title, Text: ad.Text}, nil)

	created, err := svc.Create(ad)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), created.ID)
	assert.Equal(t, ad.Title, created.Title)
	assert.Equal(t, ad.Text, created.Text)

	mockRepo.AssertExpectations(t)
}

func TestNegative_Update_ValidationFails(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	ad := ads.Ad{ID: 1, Title: "", Text: "text"}
	_, err := svc.Update(ad)
	assert.ErrorIs(t, err, ErrValidationEmptyTitle)
}

func TestPositive_Update_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	ad := ads.Ad{ID: 1, Title: "valid", Text: "text"}

	mockRepo.On("Replace", mock.MatchedBy(func(a ads.Ad) bool {
		return a.ID == ad.ID && a.Title == ad.Title && a.Text == ad.Text
	})).Return(ad, nil)

	updated, err := svc.Update(ad)
	assert.NoError(t, err)
	assert.Equal(t, ad, updated)

	mockRepo.AssertExpectations(t)
}

func TestPositive_GetById_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	ad := ads.Ad{ID: 1, Title: "title"}

	mockRepo.On("GetAdById", int64(1)).Return(ad, nil)

	got, err := svc.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, ad, got)

	mockRepo.AssertExpectations(t)
}

func TestNegative_GetById_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	mockRepo.On("GetAdById", int64(1)).Return(ads.Ad{}, adrepo.ErrNotFound)

	_, err := svc.GetById(1)
	assert.ErrorIs(t, err, ErrAdNotFound)

	mockRepo.AssertExpectations(t)
}

func TestPositive_List(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := New(mockRepo)

	list := []ads.Ad{
		{ID: 1, Title: "a"},
		{ID: 2, Title: "b"},
	}
	filter := ads.NewFilter()
	mockRepo.On("List", filter).Return(list, nil)

	got, err := svc.List(filter)
	assert.NoError(t, err)
	assert.Equal(t, list, got)

	mockRepo.AssertExpectations(t)
}
