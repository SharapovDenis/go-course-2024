package usecase

import (
	"homework8/internal/models/ads"
	adsvc "homework8/internal/services/ad"
)

type ListAdUseCase struct {
	ad adsvc.Service
}

func ListAd(adSvc adsvc.Service) *ListAdUseCase {
	return &ListAdUseCase{ad: adSvc}
}

func (u *ListAdUseCase) Execute(f *ads.Filter) ([]ads.Ad, error) {
	return u.ad.List(f)
}
