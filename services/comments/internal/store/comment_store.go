package store

import (
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/types"
)

type CommentStore struct {
	mu       sync.RWMutex
	comments []types.Comment
	counter  atomic.Int64
}

func NewCommentStore() *CommentStore {
	return &CommentStore{}
}

func (s *CommentStore) Add(userID, postID int64, body string) types.Comment {
	c := types.Comment{
		ID:        s.counter.Add(1),
		UserID:    userID,
		PostID:    postID,
		Body:      body,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	s.mu.Lock()
	s.comments = append(s.comments, c)
	s.mu.Unlock()

	return c
}

func (s *CommentStore) ListByPostID(postID int64) []types.Comment {
	s.mu.RLock()

	var result []types.Comment
	for _, c := range s.comments {
		if c.PostID == postID {
			result = append(result, c)
		}
	}

	s.mu.RUnlock()

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt > result[j].CreatedAt
	})

	return result
}
