package server

import (
	"context"

	"github.com/Feinot/authorization-service/inteenal/config"
	"github.com/Feinot/authorization-service/inteenal/handlers"
	"github.com/Feinot/authorization-service/inteenal/logger"
	"github.com/Feinot/authorization-service/inteenal/tokens"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	config, err := config.NewConfig("C:/Users/Alex/Desktop/avtor/inteenal/config/conf.yaml")
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
