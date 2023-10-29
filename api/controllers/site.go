package controllers

import (
	ae "api/errors"
	"api/models"
	"api/presenters"
	"api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type siteController struct {
	siteUsecase   usecases.SiteUsecase
	sitePresenter presenters.SitePresenter
}

type SiteController interface {
	View(*gin.Context)
	ViewAny(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Verify(*gin.Context)
}

func (c *siteController) View(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	site, err := c.siteUsecase.GetSite(id)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.sitePresenter.View(site))
}

func (c *siteController) ViewAny(ctx *gin.Context) {
	filters := map[string]string{
		"search": ctx.Query("search"),
	}

	sites, err := c.siteUsecase.GetSites(filters)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.sitePresenter.ViewAny(sites))
}

func (c *siteController) Create(ctx *gin.Context) {
	var siteCreateRequest models.SiteCreateRequest
	err := ctx.ShouldBindJSON(&siteCreateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	site, err := c.siteUsecase.CreateSite(siteCreateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.sitePresenter.View(site))
}

func (c *siteController) Update(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	var siteUpdateRequest models.SiteUpdateRequest
	err := ctx.ShouldBindJSON(&siteUpdateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	site, err := c.siteUsecase.UpdateSite(id, siteUpdateRequest)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.sitePresenter.View(site))
}

func (c *siteController) Delete(ctx *gin.Context) {
	id := getIdentifierFromRequest("id", ctx)

	site, err := c.siteUsecase.DeleteSite(id)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, c.sitePresenter.View(site))
}

func (c *siteController) Verify(ctx *gin.Context) {
	// TODO: start job verification
}

func NewSiteController(uc usecases.SiteUsecase, up presenters.SitePresenter) SiteController {
	return &siteController{
		siteUsecase:   uc,
		sitePresenter: up,
	}
}
