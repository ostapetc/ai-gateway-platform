package store

import (
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/ostapetc/ai-gateway-platform/services/users/internal/types"
)

type UserStore struct {
	mu      sync.RWMutex
	users   []types.User
	counter atomic.Uint64
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) Add(username, email string) types.User {
	u := types.User{
		ID:       s.counter.Add(1),
		Username: username,
		Email:    email,
	}

	s.mu.Lock()
	s.users = append(s.users, u)
	s.mu.Unlock()

	return u
}

func (s *UserStore) Get(id uint64) (types.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, u := range s.users {
		if u.ID == id {
			return u, true
		}
	}

	return types.User{}, false
}

func (s *UserStore) GetRandom() (types.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.users) == 0 {
		return types.User{}, false
	}

	return s.users[rand.Intn(len(s.users))], true
}
