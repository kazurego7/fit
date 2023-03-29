package stash

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var SnapCmd = &cobra.Command{
	Use:   "snap",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := stashPush()
		if exitCode != 0 {
			return
		}
		stashApply()
	},
}

func stashPush() int {
	gitSubCmd := []string{"stash", "push", "--include-untracked"}
	fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}

func stashApply() int {
	gitSubCmd := []string{"stash", "apply", "--index", "--quiet"}
	fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}
