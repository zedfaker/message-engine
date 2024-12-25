package run

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"message-engine/entity/config"
)

var ManagerApp *ManagerApplication

type ManagerApplication struct {
	ApplicationBase
}

func ManagerRun() error {
	ManagerApp = &ManagerApplication{}
	if err := initApplication(&ManagerApp.ApplicationBase); err != nil {
		return errors.Wrap(err, "初始化失败")
	}
	return ManagerApp.engine.(*gin.Engine).Run(config.AppGlobalConfig.Server.Port)
}
