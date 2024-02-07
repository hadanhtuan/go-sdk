package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/hadanhtuan/go-sdk/config"
	"gorm.io/gorm"
)

type App struct {
	Router  *gin.Engine
	DBOrm   *gorm.DB
	Config  *config.Config
	Handler map[string]interface{}
}
