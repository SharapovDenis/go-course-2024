package usecase

import (
	"homework8/internal/models/ads"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_GetAdById(t *testing.T) {
	mockAdSvc := new(MockAdService)
	uc := GetAdById(mockAdSvc)

	ad := ads.Ad{ID: 1, Title: "title", AuthorID: 2}
	mockAdSvc.On("GetById", int64(1)).Return(ad, nil)

	got, err := uc.Execute(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, ad, got)

	mockAdSvc.AssertExpectations(t)
}
