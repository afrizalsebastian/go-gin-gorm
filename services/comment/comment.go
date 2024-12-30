package comment_services

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/afrizalsebastian/go-gin-gorm/repositories"
)

func toCommentResponse(username string, comment *models.Comment, post *models.Post) *dtos.CommentResponse {
	return &dtos.CommentResponse{
		ID:       int(comment.ID),
		Content:  comment.Content,
		PostId:   int(post.ID),
		Username: username,
	}
}

func Create(claims *middleware.AppClaims, postId int, commentRequest *dtos.CreateCommentRequest) (*dtos.CommentResponse, error) {
	post, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "Post Not Found")
	}

	userId := uint(claims.ID)
	comment := &models.Comment{
		Content: commentRequest.Content,
		UserId:  &userId,
		PostId:  &post.ID,
	}

	if err := repositories.CreateComment(comment); err != nil {
		return nil, err
	}

	return toCommentResponse(claims.Username, comment, post), nil
}
