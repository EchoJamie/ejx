/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package ejx

import (
	"github.com/EchoJamie/ejx/init/global"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	AddCommandGroup(CommandGroup{
		Id:    groupId,
		Title: groupTitle,
		Commands: []*cobra.Command{
			modeCmd,
		},
	})
}

var ejxCmd = &cobra.Command{
	Use:              "ejx",
	Short:            "EJX is a tool for EchoJamie",
	Long:             global.Banner,
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		InitMode()
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

func AddCommandGroup(group CommandGroup) {
	ejxCmd.AddGroup(&cobra.Group{
		ID:    group.Id,
		Title: group.Title + " Command:",
	})
	for _, command := range group.Commands {
		ejxCmd.AddCommand(command)
	}
}
