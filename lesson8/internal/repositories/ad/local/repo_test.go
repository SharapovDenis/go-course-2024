package local

import (
	"homework8/internal/models/ads"
	adrepo "homework8/internal/repositories/ad"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_CreateAndGetById(t *testing.T) {
	repo := New()

	ad := ads.New()
	ad.Title = "Test Title"
	ad.Text = "Test Text"

	created, err := repo.Create(ad)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), created.ID)

	got, err := repo.GetAdById(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, ad.Title, got.Title)
	assert.Equal(t, ad.Text, got.Text)
}

func TestNegative_GetAdById_NotFound(t *testing.T) {
	repo := New()

	_, err := repo.GetAdById(123)
	assert.ErrorIs(t, err, adrepo.ErrNotFound)
}

func TestPositive_Replace(t *testing.T) {
	repo := New()

	ad := ads.New()
	ad.Title = "Old Title"
	created, err := repo.Create(ad)
	assert.NoError(t, err)

	created.Title = "New Title"
	replaced, err := repo.Replace(created)
	assert.NoError(t, err)
	assert.Equal(t, "New Title", replaced.Title)

	got, err := repo.GetAdById(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, "New Title", got.Title)
}

func TestNegative_Replace_NotFound(t *testing.T) {
	repo := New()

	ad := ads.New()
	ad.ID = 999 // несуществующий ID
	_, err := repo.Replace(ad)
	assert.ErrorIs(t, err, adrepo.ErrNotFound)
}

func TestPositive_List(t *testing.T) {
	repo := New()

	ad1 := ads.New()
	ad1.Title = "Ad 1"
	_, err := repo.Create(ad1)
	assert.NoError(t, err)

	ad2 := ads.New()
	ad2.Title = "Ad 2"
	_, err = repo.Create(ad2)
	assert.NoError(t, err)

	list, err := repo.List(ads.NewFilter())
	assert.NoError(t, err)
	assert.Len(t, list, 2)

	titles := map[string]bool{}
	for _, ad := range list {
		titles[ad.Title] = true
	}

	assert.True(t, titles["Ad 1"], "expected 'Ad 1' in list")
	assert.True(t, titles["Ad 2"], "expected 'Ad 2' in list")
}
