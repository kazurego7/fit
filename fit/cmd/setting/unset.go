package setting

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
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
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		flag, err := scopeFlag.toGitFlag()
		if err != nil {
			return err
		}
		gitSubCmd := []string{"config", flag, "--unset", args[0]}
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
		return nil
	},
}