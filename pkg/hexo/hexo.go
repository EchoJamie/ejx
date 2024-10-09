/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"fmt"
	"github.com/EchoJamie/ejx/init/ejx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var HexoGroup = &cobra.Group{
	ID:    GroupId,
	Title: GroupTitle,
}

func init() {
	hexoCmd.CompletionOptions.DisableDefaultCmd = true
	hexoCmd.AddGroup(HexoGroup)
}

var hexoCmd = &cobra.Command{
	Short: "hexo mode is a tool for EchoJamie",
	Long:  ejx.Banner,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.CompletionOptions.DisableDefaultCmd = true
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func CheckRootPath() {
	if !viper.InConfig("hexo.root") {
		fmt.Println("请先配置hexo博客根目录, 例如: ejx init --mode hexo")
		os.Exit(1)
	}
}

func RootCmd() *cobra.Command {
	return hexoCmd
}
