package stash

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore stored stash to worktree",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var stashRevision string
		if len(args) == 0 {
			stashRevision = `stash@{0}`
		} else {
			stashRevision = args[0]
		}

		exitCode := applyKeepIndex(stashRevision)
		if exitCode != 0 {
			apply(stashRevision)
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
