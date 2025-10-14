package ad

import (
	"homework8/internal/models/ads"
	repo "homework8/internal/repositories/ad"
	"time"
)

type Service interface {
	GetById(int64) (ads.Ad, error)
	Create(ads.Ad) (ads.Ad, error)
	Update(ads.Ad) (ads.Ad, error)
	List(*ads.Filter) ([]ads.Ad, error)
}

type service struct {
	repository repo.Repository
}

func New(repo repo.Repository) Service {
	return &service{repository: repo}
}

func validateAd(ad *ads.Ad) error {

	// Название не должно быть пустым
	if ad.Title == "" {
		return ErrValidationEmptyTitle
	}

	// Название должно быть короче 100 символов
	if len(ad.Title) > 100 {
		return ErrValidationTooLongTitle
	}

	// Текст объявления не должен быть пустым
	if ad.Text == "" {
		return ErrValidationEmptyText
	}

	// Текст объявления должен быть короче 500 символов
	if len(ad.Text) > 500 {
		return ErrValidationTooLongText
	}

	return nil
}

func (s *service) GetById(id int64) (ads.Ad, error) {
	ad, err := s.repository.GetAdById(id)
	if err != nil {
		return ads.Ad{}, handleRepoError(err)
	}
	return ad, nil
}

func (s *service) Create(ad ads.Ad) (ads.Ad, error) {

	err := validateAd(&ad)
	if err != nil {
		return ads.Ad{}, err
	}

	ad.CreatedAt = time.Now()
	ad.ModifiedAt = time.Now()

	ad, err = s.repository.Create(ad)
	if err != nil {
		return ads.Ad{}, handleRepoError(err)
	}

	return ad, nil
}

func (s *service) Update(ad ads.Ad) (ads.Ad, error) {

	err := validateAd(&ad)
	if err != nil {
		return ads.Ad{}, err
	}

	ad.ModifiedAt = time.Now()

	ad, err = s.repository.Replace(ad)
	if err != nil {
		return ads.Ad{}, handleRepoError(err)
	}

	return ad, nil

}

func (s *service) List(f *ads.Filter) ([]ads.Ad, error) {
	return s.repository.List(f)
}
