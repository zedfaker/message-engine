package run

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"message-engine/cache"
	"message-engine/entity/config"
	"net/http"
	"strings"
)

// 初始化应用
func initApplication(application *ApplicationBase) error {
	if err := application.initDB(); err != nil {
		return err
	}
	application.initEngine()
	initCache, err := InitCache()
	if err != nil {
		return err
	}
	application.Cache = initCache
	return nil
}

// 初始化数据库
func (application *ApplicationBase) initDB() error {
	switch strings.ToUpper(config.AppGlobalConfig.Db.Type) {
	case "MYSQL":
		mysql, err := InitMysql(config.AppGlobalConfig.Db.Mysql.Url, config.AppGlobalConfig.Db.Mysql.UserName, config.AppGlobalConfig.Db.Mysql.Password, config.AppGlobalConfig.Db.Mysql.Database)
		if err != nil {
			return err
		}
		application.Db = mysql
		return nil
	default:
		return errors.New("不支持的数据库类型")
	}
}

// 初始化引擎
func (application *ApplicationBase) initEngine() {
	application.engine = InitWebEngine()
}

// 应用基础结构体
type ApplicationBase struct {
	Db     *gorm.DB
	engine http.Handler
	Cache  cache.Adapter
}
