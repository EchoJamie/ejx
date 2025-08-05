/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"fmt"
	"github.com/EchoJamie/ejx/pkg/ejx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	ejx.AddCommandGroup(getCommandGroup())
}

func CheckRootPath() {
	if !viper.InConfig("hexo.root") {
		fmt.Println("请先配置hexo博客根目录, 例如: ejx init --mode hexo")
		os.Exit(1)
	}
}

func getCommandGroup() ejx.CommandGroup {
	return ejx.CommandGroup{
		Id:    groupId,
		Title: groupTitle,
		Commands: []*cobra.Command{
			initCmd,
			newCmd,
			runCmd,
			publishCmd,
		},
	}
}
