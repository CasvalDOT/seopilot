package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email,max=255"`
}

type ResetPasswordRequest struct {
	Email                string `json:"email" binding:"required,email,max=255"`
	Password             string `json:"password" binding:"required,min=8,max=255,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,min=8,max=255"`
	Token                string `json:"token" binding:"required,jwt"`
}
