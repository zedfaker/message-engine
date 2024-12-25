package run

import (
	"github.com/gin-gonic/gin"
	"message-engine/entity/config"
)

func InitWebEngine() *gin.Engine {
	gin.SetMode(config.AppGlobalConfig.Server.Mode)
	return gin.New()
}
