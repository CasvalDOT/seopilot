package middlewares

import (
	ae "api/errors"
	"api/services"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication(authService services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			error := errors.New(ae.ErrorKeyAuthInvalidAuthHeader)

			err := ae.NewApiError(error.Error())
			ctx.AbortWithStatusJSON(err.Status, err)
			return
		}

		user, err := authService.GetUserFromToken(splitToken[1])
		if err != nil {
			err := ae.NewApiError(err.Error())
			ctx.AbortWithStatusJSON(err.Status, err)
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
