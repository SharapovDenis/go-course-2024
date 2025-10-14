package usecase

import (
	"fmt"
	"homework8/internal/models/ads"
	adsvc "homework8/internal/services/ad"
)

type ChangeAdStatusUseCase struct {
	ad adsvc.Service
}

func ChangeAdStatus(adSvc adsvc.Service) *ChangeAdStatusUseCase {
	return &ChangeAdStatusUseCase{ad: adSvc}
}

func (u *ChangeAdStatusUseCase) Execute(adID int64, status bool, userID int64) (ads.Ad, error) {

	// Делаем поиск объявления
	ad, err := u.ad.GetById(adID)
	if err != nil {
		return ads.New(), fmt.Errorf("%w| adId=%d, userID=%d", handleServiceError(err), adID, userID)
	}

	// Изменения доступны только автору
	if ad.AuthorID != userID {
		return ads.New(), Err4001_002
	}

	ad.Published = status
	ad, err = u.ad.Update(ad)
	if err != nil {
		return ads.New(), fmt.Errorf("%w| adId=%d, userID=%d", handleServiceError(err), adID, userID)
	}

	return ad, nil

}
