/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package ejx

import (
	"github.com/EchoJamie/ejx/init/ejx"
	"github.com/EchoJamie/ejx/pkg/hexo"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	ejxCmd.AddCommand(modeCmd)
}

var ejxCmd = &cobra.Command{
	Use:              "ejx",
	Short:            "EJX is a tool for EchoJamie",
	Long:             ejx.Banner,
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		InitMode()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.DisableFlagParsing = true
	},
	Run: func(cmd *cobra.Command, args []string) {
		var delegateCmd *cobra.Command = nil
		if args != nil && len(args) >= 1 && args[0] == "mode" {
			delegateCmd = modeCmd
		}
		switch GetModeValue() {
		case "hexo":
			delegateCmd = hexo.RootCmd()
			delegateCmd.AddCommand(modeCmd)
			delegateCmd.SetArgs(args)
			delegateCmd.Flags().AddFlagSet(cmd.Flags())
			_ = delegateCmd.Execute()
		default:
			_ = cmd.Help()
		}

	},
}

func Run() {
	ejxCmd.CompletionOptions.DisableDefaultCmd = true
	err := ejxCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
