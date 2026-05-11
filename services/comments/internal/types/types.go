package types

type Comment struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	PostID    int64  `json:"post_id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type CreateCommentRequest struct {
	UserID int64  `json:"user_id"`
	PostID int64  `json:"post_id"`
	Body   string `json:"body"`
}

type CreateCommentResponse struct {
	ID int64 `json:"id"`
}

type ListCommentsRequest struct {
	PostID int64 `path:"post_id"`
}

type ListCommentsResponse struct {
	Comments []Comment `json:"comments"`
}
