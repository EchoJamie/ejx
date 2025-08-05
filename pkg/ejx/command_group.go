package ejx

import "github.com/spf13/cobra"

// CommandGroup 命令组结构 注册时进行传递
type CommandGroup struct {
	Id       string
	Title    string
	Commands []*cobra.Command
}
