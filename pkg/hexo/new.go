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
	hexoCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "ejx new",
	Long:    "ejx -m hexo new",
	GroupID: GroupId,
	PreRun: func(cmd *cobra.Command, args []string) {
		CheckRootPath()
	},
	Run: func(cmd *cobra.Command, args []string) {
		newCmdMain(args)
	},
}

func newCmdMain(args []string) {
	fmt.Print("请输入标题名称: ")
	var title string
	length, err := fmt.Scanln(&title)
	if err != nil {
		return
	}
	if length < 1 {
		fmt.Println("输入有误")
		os.Exit(1)
	}
	err = hexoNewDraft(title)
	if err != nil {
		fmt.Println("执行失败:", err)
		os.Exit(1)
	} else {
		fmt.Println("「" + title + "」 创建成功！🎉🎉🎉")
		open()
	}
}
