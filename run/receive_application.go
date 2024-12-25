package run

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"message-engine/entity/config"
)

var ReceiveApp *ReceiveApplication

type ReceiveApplication struct {
	ApplicationBase
}

func ReceiveRun() error {
	ReceiveApp = &ReceiveApplication{}
	if err := initApplication(&ReceiveApp.ApplicationBase); err != nil {
		return errors.Wrap(err, "初始化失败")
	}
	return ReceiveApp.engine.(*gin.Engine).Run(config.AppGlobalConfig.Server.Port)
}
