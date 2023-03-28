package conflict

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"restore", "--merge", args[0]}
		allArgs := append(gitSubCmd, args...)
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.GitCommand(global.Flags.Dryrun, allArgs...)
	},
}
