/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package ejx

import (
	"github.com/EchoJamie/ejx/init/global"
	"github.com/EchoJamie/ejx/pkg/common/cmd_group"
	"github.com/EchoJamie/ejx/pkg/hexo"
	"github.com/EchoJamie/ejx/pkg/mode"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	commands := mode.GetCommands()
	// commands = append(commands, other_package.GetCommands())
	addCommandGroup(cmd_group.CoreCommandGroup(commands...))
	addCommandGroup(hexo.GetCommandGroup())
}

var ejxCmd = &cobra.Command{
	Use:              "ejx",
	Short:            "EJX is a tool for EchoJamie",
	Long:             global.Banner,
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		mode.InitMode()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.DisableFlagParsing = true
	},
}

func Run() {
	ejxCmd.CompletionOptions.DisableDefaultCmd = true
	err := ejxCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func addCommandGroup(group cmd_group.CommandGroup) {
	ejxCmd.AddGroup(&cobra.Group{
		ID:    group.GetId(),
		Title: group.GetTitle() + " Command:",
	})
	for _, command := range group.Commands {
		ejxCmd.AddCommand(command)
	}
}
