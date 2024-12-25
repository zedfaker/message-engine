package run

import (
	"fmt"
	"gorm.io/gorm"
	"message-engine/entity/config"
	syserror "message-engine/errors"
	"time"
)
import "gorm.io/driver/mysql"

func InitMysql(url, username, password, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, url, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, syserror.NewInitError("初始化mysql数据库异常", err)
	}

	// 获取底层的 sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		return nil, syserror.NewInitError("获取sql.Db对象发生了异常", err)
	}
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(config.AppGlobalConfig.Db.Mysql.MaxOpen)

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.AppGlobalConfig.Db.Mysql.MaxIdle)

	// 设置连接最大空闲时间
	sqlDB.SetConnMaxIdleTime(time.Duration(config.AppGlobalConfig.Db.Mysql.MaxIdleTime) * time.Second)

	// 设置连接最大生命周期
	sqlDB.SetConnMaxLifetime(time.Duration(config.AppGlobalConfig.Db.Mysql.MaxLifeTime) * time.Second)
	return db, nil
}
