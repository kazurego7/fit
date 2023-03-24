package branch

import (
	"github.com/kazurego7/fit/fit/fitio"
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
		{
			gitSubCmd := []string{"update-ref", "-m", "reset: Reset " + args[0] + " to " + args[1], "refs/heads/" + args[0], args[1]}
			fitio.PrintGitCommand(gitSubCmd...)
			fitio.ExecuteGit(gitSubCmd...)
		}
		{
			gitSubCmd := []string{"branch", "--unset-upstream", args[0]}
			fitio.PrintGitCommand(gitSubCmd...)
			fitio.ExecuteGit(gitSubCmd...)
		}
	},
}
