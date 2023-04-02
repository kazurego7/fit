package stash

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var StashCmd = &cobra.Command{
	Use:   "stash",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {

	StashCmd.AddCommand(DeleteCmd)
	StashCmd.AddCommand(ListCmd)
	StashCmd.AddCommand(RestoreCmd)
	StashCmd.AddCommand(ShowCmd)
	StashCmd.AddCommand(StoreCmd)
}

func Snap(stashMessage string) int {
	exitCode := stashPush(stashMessage)
	if exitCode != 0 {
		return exitCode
	}
	return stashApply()
}

func stashPush(stashMessage string) int {
	var gitSubCmd []string
	if stashMessage == "" {
		gitSubCmd = []string{"stash", "push", "--include-untracked"}
	} else {
		gitSubCmd = []string{"stash", "push", "--include-untracked", "--message", stashMessage}
	}
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}

func stashApply() int {
	gitSubCmd := []string{"stash", "apply", "--index", "--quiet"}
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}
