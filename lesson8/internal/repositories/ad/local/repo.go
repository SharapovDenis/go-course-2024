package local

import (
	"fmt"
	"homework8/internal/models/ads"
	adrepo "homework8/internal/repositories/ad"
	"sync"
	"time"
)

type RepositoryLocal struct {
	m       map[int64]ads.Ad
	counter int64
	mu      sync.RWMutex
}

func New() *RepositoryLocal {
	return &RepositoryLocal{
		m: make(map[int64]ads.Ad),
	}
}

func (r *RepositoryLocal) GetAdById(id int64) (ads.Ad, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ad, ok := r.m[id]
	if !ok {
		return ads.Ad{}, fmt.Errorf("ad: %w", adrepo.ErrNotFound)
	}
	return ad, nil
}

func (r *RepositoryLocal) Create(ad ads.Ad) (ads.Ad, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ad.ID = r.counter
	r.m[ad.ID] = ad

	r.counter++

	return ad, nil
}

func (r *RepositoryLocal) Replace(ad ads.Ad) (ads.Ad, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.m[ad.ID]; !ok {
		return ads.Ad{}, fmt.Errorf("ad: %w", adrepo.ErrNotFound)
	}
	r.m[ad.ID] = ad
	return ad, nil
}

func (r *RepositoryLocal) List(f *ads.Filter) ([]ads.Ad, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]ads.Ad, 0, len(r.m))
	for _, ad := range r.m {

		if f.Title != nil && *f.Title != ad.Title {
			continue
		}

		if f.Text != nil && *f.Text != ad.Text {
			continue
		}

		if f.AuthorID != nil && *f.AuthorID != ad.AuthorID {
			continue
		}

		if f.Published != nil && *f.Published != ad.Published {
			continue
		}

		if f.CreatedDate != nil {
			fdate, err := time.Parse(time.DateOnly, *f.CreatedDate)
			if err != nil {
				return nil, fmt.Errorf("f.CreatedDate: %w", err)
			}
			adDate := ad.CreatedAt.Format(time.DateOnly)
			if adDate != fdate.Format(time.DateOnly) {
				continue
			}
		}

		list = append(list, ad)
	}

	return list, nil
}
