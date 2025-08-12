package dto

import (
	"homework8/internal/models/ads"
	"time"
)

type CreateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

type ChangeAdStatusRequest struct {
	Published bool  `json:"published"`
	UserID    int64 `json:"user_id"`
}

type UpdateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

type AdResponse struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Text       string    `json:"text"`
	AuthorID   int64     `json:"author_id"`
	Published  bool      `json:"published"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (r *CreateAdRequest) ToAd() ads.Ad {
	ad := ads.New()
	ad.Title = r.Title
	ad.Text = r.Text
	ad.AuthorID = r.UserID
	ad.CreatedAt = time.Now()
	ad.ModifiedAt = time.Now()
	return ad
}
