package local

import (
	"fmt"
	"homework8/internal/models/user"
	userrepo "homework8/internal/repositories/user"
	"sync"
)

type RepositoryLocal struct {
	m       map[int64]user.User
	counter int64
	mu      sync.RWMutex
}

func New() *RepositoryLocal {
	return &RepositoryLocal{
		m: make(map[int64]user.User),
	}
}

func (r *RepositoryLocal) GetUserById(id int64) (user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	usr, ok := r.m[id]
	if !ok {
		return user.New(), fmt.Errorf("user: %w", userrepo.ErrNotFound)
	}
	return usr, nil
}

func (r *RepositoryLocal) Create(usr user.User) (user.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	usr.ID = r.counter
	r.m[usr.ID] = usr

	r.counter++

	return usr, nil
}

func (r *RepositoryLocal) Replace(usr user.User) (user.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.m[usr.ID]; !ok {
		return user.New(), fmt.Errorf("user: %w", userrepo.ErrNotFound)
	}
	r.m[usr.ID] = usr
	return usr, nil
}

func (r *RepositoryLocal) List() ([]user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]user.User, 0, len(r.m))
	for _, usr := range r.m {
		list = append(list, usr)
	}

	return list, nil
}
