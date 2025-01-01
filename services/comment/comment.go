package comment_services

import (
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
	var post = &models.Post{ID: uint(postId)}
	if err := repositories.GetPostById(post); err != nil {
		return nil, err
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

func GetById(postId int, commentId int) (*dtos.CommentResponse, error) {
	var post = &models.Post{ID: uint(postId)}
	if err := repositories.GetPostById(post); err != nil {
		return nil, err
	}

	var comment = &models.Comment{
		ID:     uint(commentId),
		PostId: &post.ID,
	}

	if err := repositories.GetCommentById(comment); err != nil {
		return nil, err
	}

	return toCommentResponse(comment.User.Username, comment, post), nil
}
