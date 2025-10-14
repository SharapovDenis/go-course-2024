package ad

import (
	"homework8/internal/models/ads"
)

type Repository interface {
	GetAdById(int64) (ads.Ad, error)
	Create(ads.Ad) (ads.Ad, error)
	Replace(ads.Ad) (ads.Ad, error)
	List(*ads.Filter) ([]ads.Ad, error)
}
