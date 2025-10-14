package usecase

import (
	"homework8/internal/models/ads"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_ListAd(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := ListAd(mockAdSvc)

	adsList := []ads.Ad{{ID: 1}, {ID: 2}}
	filter := ads.NewFilter()
	mockAdSvc.On("List", filter).Return(adsList, nil)

	got, err := uc.Execute(filter)
	assert.NoError(t, err)
	assert.Equal(t, adsList, got)

	mockAdSvc.AssertExpectations(t)
}
