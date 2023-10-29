package controllers

import (
	"api/presenters"
	"api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type optionController struct {
	optionUsecase   usecases.OptionUsecase
	optionPresenter presenters.OptionPresenter
}

type OptionController interface {
	ViewRoleOptions(ctx *gin.Context)
	ViewAlertOptions(ctx *gin.Context)
}

func (o *optionController) ViewRoleOptions(ctx *gin.Context) {
	options := o.optionUsecase.ExtractRoleOptions()

	ctx.JSON(http.StatusOK, o.optionPresenter.ViewAny(options))
}

func (o *optionController) ViewAlertOptions(ctx *gin.Context) {
	options := o.optionUsecase.ExtractAlertOptions()

	ctx.JSON(http.StatusOK, o.optionPresenter.ViewAny(options))
}

func NewOptionController(optionUsecase usecases.OptionUsecase, optionPresenter presenters.OptionPresenter) OptionController {
	return &optionController{
		optionUsecase:   optionUsecase,
		optionPresenter: optionPresenter,
	}
}
