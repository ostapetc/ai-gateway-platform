package store

import (
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/types"
)

type PostStore struct {
	mu      sync.RWMutex
	posts   []types.Post
	counter atomic.Uint64
}

func NewPostStore() *PostStore {
	return &PostStore{}
}

func (s *PostStore) Add(userID uint64, title, body string) types.Post {
	post := types.Post{
		ID:        s.counter.Add(1),
		UserID:    userID,
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
	}

	s.mu.Lock()
	s.posts = append(s.posts, post)
	s.mu.Unlock()

	return post
}

func (s *PostStore) GetRandom() (types.Post, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.posts) == 0 {
		return types.Post{}, false
	}

	post := s.posts[rand.Intn(len(s.posts))]

	return post, true
}

func (s *PostStore) List() []types.Post {
	s.mu.RLock()

	result := make([]types.Post, len(s.posts))
	copy(result, s.posts)

	s.mu.RUnlock()

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})

	return result
}
