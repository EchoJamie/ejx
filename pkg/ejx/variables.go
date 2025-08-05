/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package ejx

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
	ModeDf   = "df"
	ModeNone = "none"
)

// 默认模式
const defaultMode = ModeNone

const groupId = "core"
const groupTitle = "Core"

var (
	modeType = []string{ModeHexo, ModeDf, ModeNone}
)

const modeExample string = "  ejx mode -l\n  ejx mode --set hexo"
