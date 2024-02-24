package server

import (
	"context"

	"github.com/Feinot/authorization-service/internal/config"
	"github.com/Feinot/authorization-service/internal/handlers"
	"github.com/Feinot/authorization-service/internal/modules/logger"

	"github.com/Feinot/authorization-service/internal/tokens"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	config, err := config.NewConfig("./conf.yaml")
	if err != nil {
		logger.LogError("cannot load config", err)
	}
	db, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		logger.LogError("cannot create new client", err)
	}
	err = db.Connect(context.Background())
	if err != nil {
		logger.LogError("cannot db connected", err)
	}
	token := tokens.CreateToken(db, config.SecretKey)
	router := gin.Default()
	handlers.Register(router, token)
	router.Run()
}
