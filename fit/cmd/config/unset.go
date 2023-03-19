package config

import (
	"github.com/kazurego7/fit/fit/gitexec"
	"github.com/spf13/cobra"
)

var UnsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		flag, err := scopeFlag.toGitFlag()
		if err != nil {
			return err
		}
		gitSubCmd := []string{"config", flag, "--unset", args[0]}
		gitexec.Git(gitSubCmd...)
		return nil
	},
}
