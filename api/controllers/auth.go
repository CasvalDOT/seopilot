package controllers

import (
	ae "api/errors"
	"api/models"
	"api/services"
	"api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authUsecase usecases.AuthUsecase
	authService services.AuthService
}

type AuthController interface {
	Login(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}

func (a *authController) ForgotPassword(ctx *gin.Context) {
	var forgotPasswordRequest models.ForgotPasswordRequest
	err := ctx.ShouldBindJSON(&forgotPasswordRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	err = a.authUsecase.ForgotPassword(forgotPasswordRequest.Email)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (a *authController) ResetPassword(ctx *gin.Context) {
	var resetPasswordRequest models.ResetPasswordRequest
	err := ctx.ShouldBindJSON(&resetPasswordRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	// Check the token
	_, ok, err := a.authService.VerifyToken(resetPasswordRequest.Token)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	// Check ok status
	if !ok {
		err := ae.NewApiError(ae.ErrorKeyAuthInvalidToken)
		ctx.JSON(err.Status, err)
		return
	}

	err = a.authUsecase.ResetPassword(
		resetPasswordRequest.Email,
		resetPasswordRequest.Password,
	)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (a *authController) Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	token, expirationDate, err := a.authUsecase.AuthUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":           token,
		"expiration_date": expirationDate.Unix(),
	})
}

func NewAuthController(authUsecase usecases.AuthUsecase, authService services.AuthService) AuthController {
	return &authController{
		authUsecase: authUsecase,
		authService: authService,
	}
}
