package dtos

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	ID       int     `json:"id"`
	Role     string  `json:"role"`
	Fullname *string `json:"fullname,omitempty"`
	Bio      *string `json:"bio,omitempty"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Bio      string `json:"bio"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
