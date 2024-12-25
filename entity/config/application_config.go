package config

var (
	AppGlobalConfig = &Config{}
)

// Config 管理端配置文件实体
type Config struct {
	Server struct {
		// 端口
		Port string
		// gin运行模式
		Mode string
	}
	Db struct {
		Type  string
		Mysql struct {
			Url         string
			Suffix      string
			Database    string
			UserName    string
			Password    string
			MaxOpen     int
			MaxIdle     int
			MaxLifeTime int
			MaxIdleTime int
		}
	}
	Cache struct {
		Type  string
		Redis struct {
			Mode     string
			Host     string
			Password string
			Database int
		}
	}
}
