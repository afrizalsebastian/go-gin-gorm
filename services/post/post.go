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
