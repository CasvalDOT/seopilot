package controllers

import (
	ae "api/errors"
	"api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserFromRequest(ctx *gin.Context) models.User {
	user := ctx.MustGet("user").(models.User)
	return user
}

func getIdentifierFromRequest(identifier string, ctx *gin.Context) int {
	id := ctx.Param(identifier)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err := ae.NewApiError(err.Error())
		ctx.JSON(err.Status, err)
		return -1
	}

	return idInt
}
