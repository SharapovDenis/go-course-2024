package usecase

import (
	"fmt"
	"homework8/internal/models/ads"
	adsvc "homework8/internal/services/ad"
)

type GetAdByIdUseCase struct {
	ad adsvc.Service
}

func GetAdById(adSvc adsvc.Service) *GetAdByIdUseCase {
	return &GetAdByIdUseCase{ad: adSvc}
}

func (u *GetAdByIdUseCase) Execute(id int64, userID int64) (ads.Ad, error) {
	found, err := u.ad.GetById(id)
	if err != nil {
		return ads.New(), fmt.Errorf("%w| id=%d, userID=%d", handleServiceError(err), id, userID)
	}
	return found, nil
}
