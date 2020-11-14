package request

type RegisterUserRequest struct {
	Name     *string `json:"name"`
	Email    string  `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	UserId       string `json:"userId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"RefreshToken"`
}
