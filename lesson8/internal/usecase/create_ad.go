package usecase

import (
	"fmt"
	"homework8/internal/models/ads"
	adsvc "homework8/internal/services/ad"
	usersvc "homework8/internal/services/user"
)

type CreateAdUseCase struct {
	ad   adsvc.Service
	user usersvc.Service
}

func CreateAd(adSvc adsvc.Service, userSvc usersvc.Service) *CreateAdUseCase {
	return &CreateAdUseCase{ad: adSvc, user: userSvc}
}

func (u *CreateAdUseCase) Execute(ad ads.Ad, userID int64) (ads.Ad, error) {

	// Делаем поиск пользователя
	_, err := u.user.GetById(userID)
	if err != nil {
		return ads.New(), fmt.Errorf("%w| ad=%+v, userID=%d", handleServiceError(err), ad, userID)
	}

	created, err := u.ad.Create(ad)
	if err != nil {
		return ads.New(), fmt.Errorf("%w| ad=%+v, userID=%d", handleServiceError(err), ad, userID)
	}
	return created, nil
}
