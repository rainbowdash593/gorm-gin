package users

type CreateUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
