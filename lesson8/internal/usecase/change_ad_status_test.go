package usecase

import (
	"homework8/internal/models/ads"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPositive_ChangeAdStatus(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := ChangeAdStatus(mockAdSvc)

	ad := ads.Ad{ID: 1, Published: false, AuthorID: 2}
	updatedAd := ad
	updatedAd.Published = true

	mockAdSvc.On("GetById", int64(1)).Return(ad, nil)
	mockAdSvc.On("Update", mock.MatchedBy(func(ad ads.Ad) bool {
		return ad.ID == 1 && ad.Published == true
	})).Return(updatedAd, nil)

	got, err := uc.Execute(1, true, 2)
	assert.NoError(t, err)
	assert.Equal(t, updatedAd, got)

	mockAdSvc.AssertExpectations(t)
}

func TestNegative_ChangeAdStatus_Forbidden(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := ChangeAdStatus(mockAdSvc)

	ad := ads.Ad{ID: 1, Published: false, AuthorID: 2}

	mockAdSvc.On("GetById", int64(1)).Return(ad, nil)

	_, err := uc.Execute(1, true, 3) // userID != authorID
	assert.ErrorIs(t, err, Err4001_002)

	mockAdSvc.AssertExpectations(t)
}
