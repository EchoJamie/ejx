/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package mode

/*
mode 相关变量
*/
var mode string

var setFlag bool
var listFlag bool

// 配置变量名称
const modeStr = "mode"

// 模式类型
const (
	ModeHexo = "hexo"
	ModeGit  = "git"
	ModeNone = "none"
)

// 默认模式
const defaultMode = ModeNone

var (
	modeType = []string{ModeHexo, ModeGit, ModeNone}
)

const modeExample string = "  ejx mode -l\n  ejx mode --set hexo"
