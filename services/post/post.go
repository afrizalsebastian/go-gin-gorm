package post_services

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/afrizalsebastian/go-gin-gorm/repositories"
)

func toPostResponse(post *models.Post, user *models.User) *dtos.PostResponse {
	var username *string
	var fullname *string
	if user != nil {
		username = &user.Username
		fullname = &user.Profile.Fullname
	}

	return &dtos.PostResponse{
		ID:       int(post.ID),
		Title:    string(post.Title),
		Content:  string(post.Content),
		Username: username,
		Fullname: fullname,
	}
}

func Create(userId int, postRequest *dtos.CreatePostRequest) (*dtos.PostResponse, error) {
	user, err := repositories.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This post not found.")
	}

	post := &models.Post{
		UserId:  &user.ID,
		Title:   postRequest.Title,
		Content: postRequest.Content,
	}

	if err := repositories.CreatePost(post); err != nil {
		return nil, err
	}

	return toPostResponse(post, user), nil
}

func Get(rows int, page int) (*dtos.ListPostResponse, error) {
	offset := (page - 1) * rows

	postsQuery, err := repositories.GetPost(rows, offset)
	if err != nil {
		return nil, err
	}

	posts := []*dtos.PostResponse{}
	for _, post := range postsQuery {
		response := toPostResponse(post, post.User)
		posts = append(posts, response)
	}

	count, err := repositories.GetCountPost()
	if err != nil {
		return nil, err
	}

	totalPage := (*count + int64(rows) - 1) / int64(rows)

	result := &dtos.ListPostResponse{
		Posts:     posts,
		Page:      page,
		TotalPage: int(totalPage),
	}

	return result, nil
}

func GetById(postId int) (*dtos.PostResponse, error) {
	post, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This post not found.")
	}

	return toPostResponse(post, post.User), nil
}

func Update(postId int, updateRequest *dtos.UpdatePostRequest) (*dtos.PostResponse, error) {
	extPost, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	if extPost == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This post not found.")
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

	return toPostResponse(extPost, extPost.User), nil
}

func Delete(postId int) (*dtos.PostResponse, error) {
	post, err := repositories.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This post not found.")
	}

	if err := repositories.DeletePost(post); err != nil {
		return nil, err
	}

	return toPostResponse(post, post.User), nil
}
