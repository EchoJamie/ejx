/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package ejx

import (
	"github.com/EchoJamie/ejx/pkg/conf"
	"github.com/EchoJamie/ejx/tools/filepath"
)

// 配置文件名
// 需要在 ${project.home}/configs/ 目录下创建 ejx.default.yaml 文件
const configName = "ejx.yaml"

func init() {
	conf.InitConfig(configName, configPath())
}

// 配置存储路径
func configPath() string {
	return filepath.Expend("~/.ejx/conf")
}
