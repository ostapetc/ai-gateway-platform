package types

type Comment struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	PostID    uint64 `json:"post_id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type CreateCommentRequest struct {
	UserID uint64 `json:"user_id"`
	PostID uint64 `json:"post_id"`
	Body   string `json:"body"`
}

type CreateCommentResponse struct {
	ID uint64 `json:"id"`
}

type ListCommentsRequest struct {
	PostID uint64 `form:"post_id,optional"`
}

type ListCommentsResponse struct {
	Comments []Comment `json:"comments"`
}
