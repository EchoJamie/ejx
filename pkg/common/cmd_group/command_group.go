package cmd_group

import "github.com/spf13/cobra"

// CoreGroupId 核心命令组ID
const CoreGroupId = "core"

// CoreGroupTitle 核心命令组标题
const CoreGroupTitle = "Core"

// CommandGroup 命令组结构 注册时进行传递
type CommandGroup struct {
	id       string
	title    string
	Commands []*cobra.Command
}

func (cg CommandGroup) GetId() string {
	return cg.id
}

func (cg CommandGroup) GetTitle() string {
	return cg.title
}

func CoreCommandGroup(commands ...*cobra.Command) CommandGroup {
	return CustomCommandGroup(CoreGroupId, CoreGroupTitle, commands...)
}

func CustomCommandGroup(id, title string, commands ...*cobra.Command) CommandGroup {
	return CommandGroup{
		id:       id,
		title:    title,
		Commands: commands,
	}
}
