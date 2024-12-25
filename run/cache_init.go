package run

import (
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"message-engine/cache"
	"message-engine/entity/config"
	"strings"
)

func InitCache() (cache.Adapter, error) {
	switch strings.ToUpper(config.AppGlobalConfig.Cache.Type) {
	case "REDIS":
		initRedis, err := InitRedis()
		if err != nil {
			return nil, err
		}
		return &initRedis, err
	default:
		return nil, errors.New("不支持的cache类型")
	}
}

func InitRedis() (cache.Redis, error) {
	switch strings.ToUpper(config.AppGlobalConfig.Cache.Redis.Mode) {
	case "CLUSTER":
		rdbCluster := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(config.AppGlobalConfig.Cache.Redis.Host, ","),
			Password: config.AppGlobalConfig.Cache.Redis.Password,
		})
		if err := rdbCluster.Ping(context.Background()).Err(); err != nil {
			return cache.Redis{}, err
		}
		return cache.Redis{Client: rdbCluster}, nil
	case "SINGLE":
		rdbSingle := redis.NewClient(&redis.Options{
			Addr:     config.AppGlobalConfig.Cache.Redis.Host,
			Password: config.AppGlobalConfig.Cache.Redis.Password,
			DB:       config.AppGlobalConfig.Cache.Redis.Database,
		})
		if err := rdbSingle.Ping(context.Background()).Err(); err != nil {
			return cache.Redis{}, err
		}
		return cache.Redis{Client: rdbSingle}, nil
	default:
		return cache.Redis{}, errors.New("不支持的redis连接mode值")
	}
}
