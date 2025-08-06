/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"fmt"
	"github.com/EchoJamie/ejx/pkg/common/cmd_group"
	"github.com/spf13/viper"
	"os"
)

func CheckRootPath() {
	if !viper.InConfig("hexo.root") {
		fmt.Println("请先配置hexo博客根目录, 例如: ejx init --mode hexo")
		os.Exit(1)
	}
}

func GetCommandGroup() cmd_group.CommandGroup {
	return cmd_group.CustomCommandGroup(
		groupId, groupTitle,
		initCmd,
		newCmd,
		runCmd,
		publishCmd,
	)
}
