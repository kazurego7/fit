package config

import (
	"github.com/kazurego7/fit/fit/gitexec"
	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		targetScope := getTargetScope(flags)
		gitSubCmd := []string{"config", targetScope, "--edit"}
		gitexec.Git(gitSubCmd...)
	},
}
