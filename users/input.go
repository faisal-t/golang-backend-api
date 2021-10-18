package users

type RegisterUserInput struct {
	Name         string `json:"name" binding:"required"`
	Occupation   string `json:"occupation" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	PasswordHash string `json:"password" binding:"required"`
}
