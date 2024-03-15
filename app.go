package sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/config"
)

type App struct {
	Router  *gin.Engine
	Config  *config.Config
	Handler map[string]interface{}
}
