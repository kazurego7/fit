package conflict

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var AbortCmd = &cobra.Command{
	Use:   "abort",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"merge", "--abort"}
		allArgs := append(gitSubCmd, args...)
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.CommandGit(global.Flags.Dryrun, allArgs...)
	},
}
