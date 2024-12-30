package dtos

type CommentResponse struct {
	ID       int    `json:"id"`
	PostId   int    `json:"post_id"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

type CreateCommentRequest struct {
	Content string `json:"content"`
}
