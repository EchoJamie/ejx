/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"fmt"
	"github.com/EchoJamie/ejx/tools/filepath"
	"github.com/EchoJamie/ejx/tools/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	hexoCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "ejx init",
	Long:    "ejx -m hexo init",
	GroupID: GroupId,
	Run: func(cmd *cobra.Command, args []string) {
		initCmdMain(args)
	},
}

func initCmdMain(args []string) {
	hexoRootConf()
}

func hexoRootConf() {
	var blogPath string
	hexoRootPath := viper.GetString("hexo.root")
	fmt.Println("请输入hexo博客根目录: ", text.WhiteText(hexoRootPath))
	inputLength, _ := fmt.Scanln(&blogPath)
	if inputLength <= 0 || !filepath.Exist(blogPath) {
		fmt.Println("目录不存在: ", blogPath)
		return
	}
	fmt.Println("hexo博客根目录:", blogPath)
	viper.Set("hexo.root", filepath.Expend(blogPath))
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("写入配置文件失败:", err)
	} else {
		fmt.Println("写入配置文件成功")
	}
}
