package pkg

import (
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router  *gin.Engine
	Config  *config.Config
	Handler map[string]interface{}
}
