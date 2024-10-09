/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package ejx

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	modeCmd.PersistentFlags().BoolVar(&setFlag, "set", false, "set mode")
	modeCmd.PersistentFlags().BoolVarP(&showFlag, "show", "v", false, "show mode")
}

func InitMode() {
	// 读取Mode配置
	if mode == "" {
		mode = viper.GetString(modeStr)
	}
	if !viper.InConfig(modeStr) || mode != viper.GetString(modeStr) {
		if setMode(mode) != nil {
			panic("初始化模式失败")
		}
	}
}

func GetModeValue() string {
	return mode
}

func setMode(modeValue string) error {
	containsMode(modeValue)
	viper.Set(modeStr, modeValue)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	fmt.Println("已切换至「", modeValue, "」模式")
	return nil
}

// 声明一个常量数组
var modeType = []string{"hexo"}

func containsMode(modeValue string) bool {
	if modeValue == "" {
		fmt.Println("请指定模式: ejx -m [hexo]")
		os.Exit(1)
	}

	// 判断是否包含指定模式
	for _, v := range modeType {
		if modeValue == v {
			return true
		}
	}
	fmt.Println("未知模式:", modeValue)
	return false
}

var modeCmd = &cobra.Command{
	Use:   "mode",
	Short: "设置模式",
	Run: func(cmd *cobra.Command, args []string) {
		if setFlag {
			err := setMode(args[0])
			if err != nil {
				fmt.Println("设置模式失败, 请检查参数是否正确")
			}
			return
		}
		if showFlag {
			fmt.Println("当前模式:", mode)
		}
		if !showFlag && !setFlag {
			_ = cmd.Help()
		}
	},
}
