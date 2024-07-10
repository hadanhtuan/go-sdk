package sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/config"
)

type App struct {
	Router      *gin.Engine
	Config      *config.Config
	Handler     map[string]interface{}
	CronJobList []*CronJob
}
// TODO: should add amqp, grpc connection, databases connection into App

// Create App.Start() func that run the app
