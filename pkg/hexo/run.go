/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"fmt"
	"github.com/EchoJamie/ejx/pkg/ejx"
	"github.com/spf13/cobra"
	"os"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "本地运行",
	GroupID: groupId,
	PreRun: func(cmd *cobra.Command, args []string) {
		ejx.CheckCurrentMode(ejx.ModeHexo)
		CheckRootPath()
	},
	Run: func(cmd *cobra.Command, args []string) {
		runCmdMain(args)
	},
}

func runCmdMain(args []string) {
	err := hexoClean()
	if err == nil {
		err = hexoGenerate(false)
		if err == nil {
			err = hexoServer(false)
		}
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
