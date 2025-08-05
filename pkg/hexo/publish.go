/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"fmt"
	"github.com/EchoJamie/ejx/pkg/ejx"
	"github.com/EchoJamie/ejx/tools/filepath"
	"github.com/EchoJamie/ejx/tools/interaction"
	"github.com/EchoJamie/ejx/tools/text"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const draftsSubPath = "/source/_drafts"

var publishCmd = &cobra.Command{
	Use:     "publish",
	Short:   "发布草稿",
	GroupID: groupId,
	PreRun: func(cmd *cobra.Command, args []string) {
		ejx.CheckCurrentMode(ejx.ModeHexo)
		CheckRootPath()
	},
	Run: func(cmd *cobra.Command, args []string) {
		publishCmdMain(args)
	},
}

func publishCmdMain(args []string) {
	draftsPath := extendsBlogRootPath(draftsSubPath)
	if filepath.Exist(draftsPath) {
		dir, _ := os.ReadDir(draftsPath)
		for i, file := range dir {
			fmt.Println(i+1, file.Name())
		}
		fmt.Print("请选择要发布的草稿文件序号: ")
		var index int
		_, _ = fmt.Scanln(&index)
		if index < 1 || index > len(dir) {
			fmt.Println("输入有误")
			os.Exit(1)
		}
		result := dir[index-1].Name()
		result = strings.Split(result, ".")[0]
		// 交互 - 二次确认
		interaction.SecondCheck(fmt.Sprintf("是否确认发布草稿文件(%s)?", text.WhiteText(result)))
		err := hexoPublish(result)
		if err != nil {
			fmt.Println("执行失败:", err)
			os.Exit(1)
		} else {
			fmt.Println(result + " 发布成功！🎉🎉🎉")
		}
	}
}
