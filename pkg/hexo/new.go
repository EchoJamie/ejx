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

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "新建草稿",
	GroupID: groupId,
	PreRun: func(cmd *cobra.Command, args []string) {
		ejx.CheckCurrentMode(ejx.ModeHexo)
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
		_ = open()
	}
}
