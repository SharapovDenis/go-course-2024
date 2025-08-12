// Здесь делаем адаптацию фреймворка к общему dto

package httpgin

import (
	"homework8/internal/models/ads"
	"homework8/internal/ports/dto"

	"github.com/gin-gonic/gin"
)

func AdSuccessResponse(ad *ads.Ad) gin.H {
	return gin.H{
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

func AdErrorResponse(err error) gin.H {
	return gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func ListAdResponse(adList []ads.Ad) gin.H {
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
	return gin.H{
		"data": adResponses,
	}
}
