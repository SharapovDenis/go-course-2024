package usecase

import (
	"homework8/internal/models/ads"
	adsvc "homework8/internal/services/ad"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPositive_UpdateAd(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := UpdateAd(mockAdSvc)

	origAd := ads.Ad{ID: 1, Title: "old", Text: "old text", AuthorID: 2}
	updatedAd := origAd
	updatedAd.Title = "new"
	updatedAd.Text = "new text"

	mockAdSvc.On("GetById", int64(1)).Return(origAd, nil)
	mockAdSvc.On("Update", mock.MatchedBy(func(ad ads.Ad) bool {
		return ad.ID == 1 && ad.Title == "new" && ad.Text == "new text"
	})).Return(updatedAd, nil)

	got, err := uc.Execute(1, "new", "new text", 2)
	assert.NoError(t, err)
	assert.Equal(t, updatedAd, got)

	mockAdSvc.AssertExpectations(t)
}

func TestNegative_UpdateAd_Forbidden(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := UpdateAd(mockAdSvc)

	origAd := ads.Ad{ID: 1, Title: "old", AuthorID: 2}

	mockAdSvc.On("GetById", int64(1)).Return(origAd, nil)

	_, err := uc.Execute(1, "new", "new text", 3) // userID != authorID
	assert.ErrorIs(t, err, Err4001_002)

	mockAdSvc.AssertExpectations(t)
}

func TestNegative_UpdateAd_NotFound(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := UpdateAd(mockAdSvc)

	mockAdSvc.On("GetById", int64(1)).Return(ads.Ad{}, adsvc.ErrAdNotFound)

	_, err := uc.Execute(1, "new", "new text", 2)
	assert.ErrorIs(t, err, Err4002_001)

	mockAdSvc.AssertExpectations(t)
}
