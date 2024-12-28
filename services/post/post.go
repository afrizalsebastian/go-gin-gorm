package post_services

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/afrizalsebastian/go-gin-gorm/repositories"
)

func toPostRequest(post *models.Post, user *models.User) *dtos.PostResponse {
	return &dtos.PostResponse{
		ID:       int(post.ID),
		Title:    string(post.Title),
		Content:  string(post.Content),
		Username: user.Username,
		Fullname: user.Profile.Fullname,
	}
}

func Create(userId int, postRequest *dtos.CreatePostRequest) (*dtos.PostResponse, error) {
	user, err := repositories.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This user not found.")
	}

	post := &models.Post{
		UserId:  &user.ID,
		Title:   postRequest.Title,
		Content: postRequest.Content,
	}

	if err := repositories.CreatePost(post); err != nil {
		return nil, err
	}

	return toPostRequest(post, user), nil
}

func GetById(postId int) (*dtos.PostResponse, error) {
	post, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This user not found.")
	}

	return toPostRequest(post, post.User), nil
}

func Update(postId int, updateRequest *dtos.UpdatePostRequest) (*dtos.PostResponse, error) {
	extPost, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	if extPost == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This user not found.")
	}

	if updateRequest.Content != nil {
		extPost.Content = *updateRequest.Content
	}

	if updateRequest.Title != nil {
		extPost.Title = *updateRequest.Title
	}

	if err := repositories.UpdatePost(extPost); err != nil {
		return nil, err
	}

	return toPostRequest(extPost, extPost.User), nil
}
