package handlers

import (
	"github.com/Feinot/authorization-service/internal/entity"
	"github.com/Feinot/authorization-service/internal/errors"
	"github.com/Feinot/authorization-service/internal/storage"
	"github.com/gin-gonic/gin"
)

const (
	unauthorizedCode        = 401
	internalServerErrorCode = 500
)

func RefreshMiddleware(next gin.HandlerFunc, t *entity.Tokens) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken := c.Param("refresh")
		guid := c.Param("guid")
		if refreshToken == "" {
			errors.RespondWithError(c, unauthorizedCode, "Invalid API token")
			return
		}
		ok, err := storage.CheckDb(refreshToken, guid, t)
		if err != nil {

			errors.RespondWithError(c, internalServerErrorCode, "Server error")

			return
		}
		if ok {
			next(c)
		} else {
			errors.RespondWithError(c, unauthorizedCode, "Invalid API token")
			return
		}
	}
}
