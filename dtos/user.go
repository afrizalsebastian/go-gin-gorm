package dtos

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Role     string `json:"role"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
