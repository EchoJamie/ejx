/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	hexoCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "ejx run",
	Long:    "ejx -m hexo run",
	GroupID: GroupId,
	PreRun: func(cmd *cobra.Command, args []string) {
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
