package stash

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := applyKeepIndex(args[0])
		if exitCode != 0 {
			apply(args[0])
		}
	},
}

func applyKeepIndex(stashRevision string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", "--index", stashRevision}
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
}

func apply(stashRevision string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", stashRevision}
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	return util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
}
