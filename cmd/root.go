package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"message-engine/pkg"
)

var (
	rootCmd = &cobra.Command{
		// 预处理
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
)

func init() {

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("执行cli命令发生了异常")
		panic(err)
	}
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`管理系统`) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}
