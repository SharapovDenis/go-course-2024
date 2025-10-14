package usecase

import (
	"homework8/internal/models/ads"
	"homework8/internal/models/user"
	adsvc "homework8/internal/services/ad"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_CreateAd(t *testing.T) {
	mockAdSvc := new(MockAdService)
	mockUserSvc := new(MockUserService)
	uc := CreateAd(mockAdSvc, mockUserSvc)

	usr := user.User{ID: 2, Name: "Alice"}

	ad := ads.Ad{Title: "title", Text: "text", AuthorID: usr.ID}

	created := ad
	created.ID = 1

	mockUserSvc.On("GetById", usr.ID).Return(usr, nil)
	mockAdSvc.On("Create", ad).Return(created, nil)

	got, err := uc.Execute(ad, usr.ID)

	assert.NoError(t, err)
	assert.Equal(t, created, got)

	mockUserSvc.AssertExpectations(t)
	mockAdSvc.AssertExpectations(t)
}

func TestNegative_ChangeAdStatus_NotFound(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := ChangeAdStatus(mockAdSvc)

	mockAdSvc.On("GetById", int64(1)).Return(ads.Ad{}, adsvc.ErrAdNotFound)

	_, err := uc.Execute(1, true, 2)
	assert.ErrorIs(t, err, Err4002_001)

	mockAdSvc.AssertExpectations(t)
}
