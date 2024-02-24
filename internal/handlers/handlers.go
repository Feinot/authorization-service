package handlers

import (
	"github.com/Feinot/authorization-service/internal/entity"
	"github.com/Feinot/authorization-service/internal/errors"
	"github.com/Feinot/authorization-service/internal/modules/logger"

	"github.com/Feinot/authorization-service/internal/storage"
	"github.com/Feinot/authorization-service/internal/tokens"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, t *entity.Tokens) {
	router.GET("/login/:guid", func(c *gin.Context) {
		AuthHandlers(c, t)
	})
	router.GET("/refresh/:guid/:refresh", RefreshMiddleware(func(c *gin.Context) {
		AuthHandlers(c, t)
	}, t))
}

func AuthHandlers(c *gin.Context, t *entity.Tokens) {

	guid := c.Param("guid")

	at, u, err := tokens.GenerateToken(guid, t)

	if err != nil {
		logger.LogError("cannot generate token", err)
		errors.RespondWithError(c, internalServerErrorCode, "can`t generate token")
		return
	}

	err = storage.CreateDb(u, at, t)

	if err != nil {
		logger.LogError("cannot create db", err)
		errors.RespondWithError(c, internalServerErrorCode, "db error")
		return
	}
	c.JSON(200, at)

}
