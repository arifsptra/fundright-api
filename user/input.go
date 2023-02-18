package user

// struct for mapping user register input
type RegisterUserInput struct {
	Name           string `json:"name" binding:"required"`
	Occupation     string `json:"occupation" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
}

// struct for mapping user login input
type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}