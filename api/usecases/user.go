package usecases

import (
	"api/emails"
	ae "api/errors"
	"api/events"
	"api/models"
	"api/repositories"
	"api/services"
	"errors"
	"fmt"
	"time"
)

type userUsecase struct {
	authService  services.AuthService
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	onUserCreate events.Event
}

type UserUsecase interface {
	GetUser(int) (models.User, error)
	GetUsers(map[string]string) ([]models.User, error)
	CreateUser(models.UserCreateRequest) (models.User, error)
	UpdateUser(int, models.UserUpdateRequest) (models.User, error)
	DeleteUser(int) (models.User, error)
}

func (u *userUsecase) GetUser(id int) (models.User, error) {
	user, err := u.userRepo.GetOne(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUsecase) GetUsers(filters map[string]string) ([]models.User, error) {
	users, err := u.userRepo.GetMany(filters)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *userUsecase) CreateUser(payload models.UserCreateRequest) (models.User, error) {
	user := models.User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	// Check if email already exists
	oldUser, _ := u.userRepo.GetByEmail(user.Email)
	if oldUser.ID != 0 {
		return user, errors.New(ae.ErrorKeyUserMailAlreadyExists)
	}

	// Take the roles
	roles, _ := u.roleRepo.GetMany(map[string]interface{}{
		"ids": payload.Roles,
	})
	user.Roles = roles

	newUser, err := u.userRepo.Insert(user)
	if err != nil {
		return newUser, errors.New(ae.ErrorKeyUserCannotCreate)
	}

	token, err := u.authService.CreateToken(&user, services.Claim{
		Exp: time.Now().Add(time.Minute * 10),
		Sub: "forgot-password",
	})
	if err != nil {
		return newUser, err
	}

	// Process user to send email
	u.onUserCreate.Dispatch(emails.UserActivateAccountTemplateData{
		User: newUser,
		Url:  fmt.Sprintf("https://dashboard.seopilot.dev/reset-password?token=%s", token),
	})

	return newUser, nil
}

func (u *userUsecase) UpdateUser(id int, payload models.UserUpdateRequest) (models.User, error) {
	user, err := u.userRepo.GetOne(id)
	if err != nil || user.ID == 0 {
		return user, errors.New(ae.ErrorKeyUserNotFound)
	}

	// Check if email already exists
	userEmailOwner, _ := u.userRepo.GetByEmail(payload.Email)
	if userEmailOwner.ID != 0 && userEmailOwner.ID != user.ID {
		return user, errors.New(ae.ErrorKeyUserMailAlreadyExists)
	}

	// Check if roles are valid
	user.Name = payload.Name
	user.Email = payload.Email
	roles, _ := u.roleRepo.GetMany(map[string]interface{}{
		"ids": payload.Roles,
	})
	user.Roles = roles

	user, err = u.userRepo.Save(user)
	if err != nil {
		fmt.Println(err)
		return user, errors.New(ae.ErrorKeyUserCannotUpdate)
	}

	return user, nil
}

func (u *userUsecase) DeleteUser(id int) (models.User, error) {
	user, err := u.userRepo.GetOne(id)
	if err != nil || user.ID == 0 {
		return user, errors.New(ae.ErrorKeyUserNotFound)
	}

	user, err = u.userRepo.Delete(id)
	if err != nil {
		return user, errors.New(ae.ErrorKeyUserCannotDelete)
	}

	return user, nil
}

func NewUserUsecase(
	authService services.AuthService,
	userRepo repositories.UserRepository,
	roleRepo repositories.RoleRepository,
	onUserCreate events.Event,
) UserUsecase {
	return &userUsecase{
		authService:  authService,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		onUserCreate: onUserCreate,
	}
}
