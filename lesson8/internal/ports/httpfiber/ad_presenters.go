// Здесь делаем адаптацию фреймворка к общему dto

package httpfiber

import (
	"github.com/gofiber/fiber/v2"

	"homework8/internal/models/ads"
	"homework8/internal/ports/dto"
)

func AdSuccessResponse(ad *ads.Ad) *fiber.Map {
	return &fiber.Map{
		"data": dto.AdResponse{
			ID:         ad.ID,
			Title:      ad.Title,
			Text:       ad.Text,
			AuthorID:   ad.AuthorID,
			Published:  ad.Published,
			CreatedAt:  ad.CreatedAt,
			ModifiedAt: ad.ModifiedAt,
		},
		"error": nil,
	}
}

func AdErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  nil,
		"error": err.Error(),
	}
}

func ListAdResponse(adList []ads.Ad) *fiber.Map {
	adResponses := make([]dto.AdResponse, 0, len(adList))
	for _, ad := range adList {
		adResponses = append(adResponses, dto.AdResponse{
			ID:         ad.ID,
			Title:      ad.Title,
			Text:       ad.Text,
			AuthorID:   ad.AuthorID,
			Published:  ad.Published,
			CreatedAt:  ad.CreatedAt,
			ModifiedAt: ad.ModifiedAt,
		})
	}
	return &fiber.Map{
		"data": adResponses,
	}
}
