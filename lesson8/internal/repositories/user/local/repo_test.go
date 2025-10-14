package local

import (
	"homework8/internal/models/user"
	userrepo "homework8/internal/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositive_CreateAndGetUserById(t *testing.T) {
	repo := New()

	usr := user.New()
	usr.Name = "John Doe"
	usr.Email = "john@example.com"

	created, err := repo.Create(usr)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), created.ID)
	assert.Equal(t, "John Doe", created.Name)
	assert.Equal(t, "john@example.com", created.Email)

	got, err := repo.GetUserById(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, created, got)
}

func TestNegative_GetUserById_NotFound(t *testing.T) {
	repo := New()

	_, err := repo.GetUserById(123)
	assert.ErrorIs(t, err, userrepo.ErrNotFound)
}

func TestPositive_Replace(t *testing.T) {
	repo := New()

	usr := user.New()
	usr.Name = "Alice"
	created, err := repo.Create(usr)
	assert.NoError(t, err)

	created.Name = "Alice Updated"
	replaced, err := repo.Replace(created)
	assert.NoError(t, err)
	assert.Equal(t, "Alice Updated", replaced.Name)

	got, err := repo.GetUserById(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Alice Updated", got.Name)
}

func TestNegative_Replace_NotFound(t *testing.T) {
	repo := New()

	usr := user.New()
	usr.ID = 999 // несуществующий ID
	_, err := repo.Replace(usr)
	assert.ErrorIs(t, err, userrepo.ErrNotFound)
}

func TestPositive_List(t *testing.T) {
	repo := New()

	usr1 := user.New()
	usr1.Name = "User One"
	repo.Create(usr1)

	usr2 := user.New()
	usr2.Name = "User Two"
	repo.Create(usr2)

	list, err := repo.List()
	assert.NoError(t, err)
	assert.Len(t, list, 2)

	names := map[string]bool{}
	for _, u := range list {
		names[u.Name] = true
	}

	assert.Contains(t, names, "User One")
	assert.Contains(t, names, "User Two")
}
