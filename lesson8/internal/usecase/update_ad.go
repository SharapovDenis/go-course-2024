package usecase

import (
	"fmt"
	"homework8/internal/models/ads"
	adsvc "homework8/internal/services/ad"
)

type UpdateAdUseCase struct {
	ad adsvc.Service
}

func UpdateAd(adSvc adsvc.Service) *UpdateAdUseCase {
	return &UpdateAdUseCase{ad: adSvc}
}

func (u *UpdateAdUseCase) Execute(adID int64, title string, text string, userID int64) (ads.Ad, error) {

	// Делаем поиск объявления
	ad, err := u.ad.GetById(adID)
	if err != nil {
		return ads.New(), fmt.Errorf("%w| adID=%d, title=%s, text=%s, userID=%d", handleServiceError(err), adID, title, text, userID)
	}

	// Изменения доступны только автору
	if ad.AuthorID != userID {
		return ads.New(), Err4001_002
	}

	ad.Title = title
	ad.Text = text

	updated, err := u.ad.Update(ad)
	if err != nil {
		return ads.New(), fmt.Errorf("%w| adID=%d, title=%s, text=%s, userID=%d", handleServiceError(err), adID, title, text, userID)
	}

	return updated, nil
}
