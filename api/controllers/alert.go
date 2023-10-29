package controllers

import (
	ae "api/errors"
	"api/models"
	"api/presenters"
	"api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type alertController struct {
	alertUsecase   usecases.AlertUsecase
	alertPresenter presenters.AlertPresenter
}

type AlertController interface {
	View(*gin.Context)
	ViewAny(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func (c *alertController) View(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	alert, err := c.alertUsecase.GetAlert(id)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.alertPresenter.View(alert))
}

func (c *alertController) ViewAny(ctx *gin.Context) {
	filters := map[string]string{
		"search": ctx.Query("search"),
	}

	alerts, err := c.alertUsecase.GetAlerts(filters)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.alertPresenter.ViewAny(alerts))
}

func (c *alertController) Create(ctx *gin.Context) {
	var alertCreateRequest models.AlertCreateRequest
	err := ctx.ShouldBindJSON(&alertCreateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	alert, err := c.alertUsecase.CreateAlert(alertCreateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.alertPresenter.View(alert))
}

func (c *alertController) Update(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	var alertUpdateRequest models.AlertUpdateRequest
	err := ctx.ShouldBindJSON(&alertUpdateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	alert, err := c.alertUsecase.UpdateAlert(id, alertUpdateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.alertPresenter.View(alert))
}

func (c *alertController) Delete(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	alert, err := c.alertUsecase.DeleteAlert(id)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.alertPresenter.View(alert))
}

func NewAlertController(ac usecases.AlertUsecase, ap presenters.AlertPresenter) AlertController {
	return &alertController{
		alertUsecase:   ac,
		alertPresenter: ap,
	}
}
