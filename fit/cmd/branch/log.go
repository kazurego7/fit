package branch

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"reflog", args[0]}
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
	},
}
