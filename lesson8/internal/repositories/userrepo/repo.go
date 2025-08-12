package userrepo

import (
	"homework8/internal/models/enums"
	"homework8/internal/models/user"
	"sync"
)

type Repository interface {
	GetUserById(int64) (user.User, error)
	Create(user.User) (user.User, error)
	Replace(user.User) (user.User, error)
	List() ([]user.User, error)
}

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
		return user.New(), enums.ErrNotFound
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
		return user.New(), enums.ErrNotFound
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
