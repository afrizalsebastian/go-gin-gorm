package dtos

type PostResponse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	Fullname *string `json:"fullname"`
	Username *string `json:"username"`
}

type ListPostResponse struct {
	Posts     []*PostResponse `json:"id"`
	Page      int             `json:"page"`
	TotalPage int             `json:"total_page"`
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}
