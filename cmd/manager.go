package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"message-engine/entity/config"
	"message-engine/pkg"
	"message-engine/run"
)

var (
	// 管理端配置文件路径
	managerConfigPath string
	// 管理端命令行配置
	managerCmd = &cobra.Command{
		Use:   "manager",
		Short: "管理后台的执行命令",
		Long:  `管理后台的执行命令行`,
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig()
			err := run.ManagerRun()
			if err != nil {
				panic(fmt.Errorf("启动失败：%v", err))
			}
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			managerTip()
		},
	}
)

// init 初始化配置参数
func init() {
	managerCmd.PersistentFlags().StringVarP(&managerConfigPath, "config", "c", "doc/config/manager-application.yml", "默认使用的配置文件为：doc/config/application.yml)")
	rootCmd.AddCommand(managerCmd)
}

func managerTip() {
	usageStr := `当前启动的应用为： ` + pkg.Green(`hely牌管理系统`) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

// InitConfig 初始化配置文件
func InitConfig() {
	viper.SetConfigName("application") // name of config file (without extension)
	viper.SetConfigType("yml")         // REQUIRED if the config file does not have the extension in the name
	viper.SetConfigFile(managerConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(config.AppGlobalConfig); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
