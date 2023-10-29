package controllers

import (
	ae "api/errors"
	"api/models"
	"api/presenters"
	"api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUsecase   usecases.UserUsecase
	userPresenter presenters.UserPresenter
}

type UserController interface {
	View(*gin.Context)
	ViewAny(*gin.Context)
	ViewMe(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func (c *userController) ViewMe(ctx *gin.Context) {
	requestUser := getUserFromRequest(ctx)

	ctx.JSON(http.StatusOK, c.userPresenter.View(requestUser))
}

func (c *userController) View(ctx *gin.Context) {
	// Policy check

	id := getIdentifierFromRequest("id", ctx)

	user, err := c.userUsecase.GetUser(id)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.userPresenter.View(user))
}

func (c *userController) ViewAny(ctx *gin.Context) {
	filters := map[string]string{
		"search": ctx.Query("search"),
	}

	users, err := c.userUsecase.GetUsers(filters)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.userPresenter.ViewAny(users))
}

func (c *userController) Create(ctx *gin.Context) {
	var userCreateRequest models.UserCreateRequest
	err := ctx.ShouldBindJSON(&userCreateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	user, err := c.userUsecase.CreateUser(userCreateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.userPresenter.View(user))
}

func (c *userController) Update(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	var userUpdateRequest models.UserUpdateRequest
	err := ctx.ShouldBindJSON(&userUpdateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	user, err := c.userUsecase.UpdateUser(id, userUpdateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.userPresenter.View(user))
}

func (c *userController) Delete(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	user, err := c.userUsecase.DeleteUser(id)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.userPresenter.View(user))
}

func NewUserController(uc usecases.UserUsecase, up presenters.UserPresenter) UserController {
	return &userController{
		userUsecase:   uc,
		userPresenter: up,
	}
}
