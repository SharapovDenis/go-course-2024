// Здесь описывается бизнес-логика

package service

import (
	"homework8/internal/models/ads"
	"homework8/internal/models/enums"
	"homework8/internal/models/user"
	"homework8/internal/service/adsvc"
	"homework8/internal/service/usersvc"
)

type Service struct {
	Ad   adsvc.Service
	User usersvc.Service
}

func New(adService adsvc.Service, userService usersvc.Service) Service {
	return Service{Ad: adService, User: userService}
}

func (s *Service) GetAdById(id int64, userID int64) (ads.Ad, error) {
	return s.Ad.GetById(id)
}

func (s *Service) CreateAd(ad ads.Ad, userID int64) (ads.Ad, error) {

	// Делаем поиск пользователя
	_, err := s.User.GetById(userID)
	if err != nil {
		return ads.New(), err
	}

	return s.Ad.Create(ad)
}

func (s *Service) UpdateAd(adID int64, title string, text string, userID int64) (ads.Ad, error) {

	// Делаем поиск объявления
	ad, err := s.Ad.GetById(adID)
	if err != nil {
		return ads.New(), err
	}

	// Изменения доступны только автору
	if ad.AuthorID != userID {
		return ads.New(), enums.ErrForbidden
	}

	ad.Title = title
	ad.Text = text

	return s.Ad.Update(ad)
}

func (s *Service) ListAd(f *ads.Filter) ([]ads.Ad, error) {
	return s.Ad.List(f)
}

func (s *Service) ChangeAdStatus(adID int64, status bool, userID int64) (ads.Ad, error) {

	// Делаем поиск объявления
	ad, err := s.Ad.GetById(adID)
	if err != nil {
		return ads.New(), err
	}

	// Изменения доступны только автору
	if ad.AuthorID != userID {
		return ads.New(), enums.ErrForbidden
	}

	ad.Published = status
	ad, err = s.Ad.Update(ad)
	if err != nil {
		return ads.New(), err
	}

	return ad, nil

}

func (s *Service) GetUserById(id int64) (user.User, error) {
	return s.User.GetById(id)
}

func (s *Service) CreateUser(usr user.User) (user.User, error) {
	return s.User.Create(usr)
}

func (s *Service) UpdateUser(userID int64, name string, email string) (user.User, error) {

	// Делаем поиск пользователя
	usr, err := s.User.GetById(userID)
	if err != nil {
		return user.New(), err
	}

	usr.Name = name
	usr.Email = email
	return s.User.Update(usr)
}

func (s *Service) ListUsers() ([]user.User, error) {
	return s.User.List()
}
