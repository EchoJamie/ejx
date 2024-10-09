/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package ejx

import (
	"github.com/EchoJamie/ejx/pkg/conf"
	"github.com/EchoJamie/ejx/tools/filepath"
)

// 配置文件名
// 需要在 ${project.home}/configs/ 目录下创建 ejx.default.yaml 文件
const configName = "ejx"

func init() {
	conf.InitConfig(configName, configPath())
}

// 配置存储路径
func configPath() string {
	return filepath.Expend("~/.ejx/conf")
}

var Banner = `
                \=/,         _-===-_-====-_-===-_-==========-_-====-_
                |  @___oo   (                                         )
      /\  /\   / (___,,,}_--=                                         )
     ) /^\) ^\/ _)        =__       EJX is a tool for EchoJamie,       )
     )   /^\/   _)          (_       it can help you to manage         )
     )   _ /  / _)            (            your blog, etc.            )
 /\  )/\/ ||  | )_)            (_                                     )
<  >      |(,,) )__)             (                                   )
 ||      /    \)___)\             (_                                 )
 | \____(      )___) )___           -==-_____-=====-_____-=====-___==
  \______(_______;;; __;;;
`
