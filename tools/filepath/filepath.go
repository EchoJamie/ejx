/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package filepath

import (
	"github.com/mitchellh/go-homedir"
	"os"
)

func HomeDir() string {
	dir, _ := homedir.Dir()
	return dir
}

func Exist(path string) bool {
	// 处理 ～ 代表的 home 目录
	path = Expend(path)
	_, err := os.Stat(path)
	return err == nil
}

func Expend(path string) string {
	expand, _ := homedir.Expand(path)
	return expand
}
