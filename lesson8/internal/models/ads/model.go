package ads

import (
	"time"
)

type Ad struct {
	ID         int64
	Title      string
	Text       string
	AuthorID   int64
	Published  bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func New() Ad {
	now := time.Now()
	return Ad{CreatedAt: now, ModifiedAt: now}
}

type Filter struct {
	Title       *string
	Text        *string
	AuthorID    *int64
	Published   *bool
	CreatedDate *string // time.DateOnly
}

func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) SetTitle(t string) *Filter {
	f.Title = &t
	return f
}

func (f *Filter) SetText(t string) *Filter {
	f.Text = &t
	return f
}

func (f *Filter) SetAuthorID(id int64) *Filter {
	f.AuthorID = &id
	return f
}

func (f *Filter) SetPublished(p bool) *Filter {
	f.Published = &p
	return f
}

func (f *Filter) SetCreatedDate(s string) *Filter {
	f.CreatedDate = &s
	return f
}
