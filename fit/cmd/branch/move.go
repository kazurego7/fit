package branch

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var MoveCmd = &cobra.Command{
	Use:   "move",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := moveBranch(args[0], args[1])
		if exitCode != 0 {
			return
		}
		unsetUpstream(args[0])
	},
}

func moveBranch(branch string, revision string) int {
	gitSubCmd := []string{"branch", "--force", branch, revision}
	fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}

func unsetUpstream(branch string) int {
	gitSubCmd := []string{"branch", "--unset-upstream", branch}
	fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}
