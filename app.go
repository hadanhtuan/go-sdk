package sdk

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Router  *gin.Engine
	Config  *Config
	Handler map[string]interface{}
}
