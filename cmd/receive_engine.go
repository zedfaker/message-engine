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
	receiveConfigPath string
	// 管理端命令行配置
	receiverCmd = &cobra.Command{
		Use:   "receive",
		Short: "接收服务执行命令",
		Long:  `接收服务的执行命令行`,
		Run: func(cmd *cobra.Command, args []string) {
			InitConfigReceive()
			err := run.ReceiveRun()
			if err != nil {
				panic(fmt.Errorf("启动失败：%v", err))
			}
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			receiveTip()
		},
	}
)

// init 初始化配置参数
func init() {
	receiverCmd.PersistentFlags().StringVarP(&receiveConfigPath, "config", "c", "doc/config/receive-application.yml", "默认使用的配置文件为：doc/config/application.yml)")
	rootCmd.AddCommand(receiverCmd)
}

func receiveTip() {
	usageStr := `当前启动的应用为： ` + pkg.Green(`hely牌接收服务系统`) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

// InitConfig 初始化配置文件
func InitConfigReceive() {
	viper.SetConfigName("application") // name of config file (without extension)
	viper.SetConfigType("yml")         // REQUIRED if the config file does not have the extension in the name
	viper.SetConfigFile(receiveConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(config.AppGlobalConfig); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
