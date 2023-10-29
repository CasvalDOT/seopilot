package usecases

import (
	"api/emails"
	ae "api/errors"
	"api/events"
	"api/repositories"
	"api/services"
	"errors"
	"fmt"
	"time"
)

type authUsecase struct {
	authService      services.AuthService
	userRepo         repositories.UserRepository
	onForgotPassword events.Event
}

type AuthUsecase interface {
	AuthUser(email, password string) (string, time.Time, error)
	ForgotPassword(email string) error
	ResetPassword(email, password string) error
}

func (u *authUsecase) ResetPassword(email, password string) error {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return errors.New(ae.ErrorKeyUserNotFound)
	}

	user.Password, err = u.authService.GeneratePasswordHash(password)
	if err != nil {
		return err
	}

	_, err = u.userRepo.Save(user)
	if err != nil {
		return errors.New(ae.ErrorKeyUserCannotUpdate)
	}

	return nil
}

func (u *authUsecase) ForgotPassword(email string) error {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return nil
	}

	// Generate token
	token, err := u.authService.CreateToken(&user, services.Claim{
		Exp: time.Now().Add(time.Minute * 10),
		Sub: "forgot-password",
	})
	if err != nil {
		return err
	}

	u.onForgotPassword.Dispatch(emails.UserForgotPasswordTemplateData{
		Email:     email,
		Name:      user.Name,
		Url:       fmt.Sprintf("https://dashboard.seopilot.dev/reset-password?token=%s", token),
		CreatedAt: time.Now(),
	})

	return nil
}

func (u *authUsecase) AuthUser(email, password string) (string, time.Time, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", time.Now(), errors.New(ae.ErrorKeyAuthEmailOrPasswordIncorrent)
	}

	isEqual, err := u.authService.VerifyPasswordHash(user.Password, password)
	if !isEqual || err != nil {
		return "", time.Now(), errors.New(ae.ErrorKeyAuthEmailOrPasswordIncorrent)
	}

	expirationDate := time.Now().Add(time.Minute * 30)

	token, err := u.authService.CreateToken(&user, services.Claim{
		Exp: expirationDate,
		Sub: "authentication",
	})

	return token, expirationDate, nil
}

func NewAuthUsecase(authService services.AuthService, userRepo repositories.UserRepository, onForgotPassword events.Event) AuthUsecase {
	return &authUsecase{
		authService:      authService,
		userRepo:         userRepo,
		onForgotPassword: onForgotPassword,
	}
}
