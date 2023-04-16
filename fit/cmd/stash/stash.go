package stash

import (
	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var StashCmd = &cobra.Command{
	Use:   "stash",
	Short: "スタッシュに関する操作.",
}

func init() {
	StashCmd.AddCommand(ListCmd)
	StashCmd.AddCommand(StoreCmd)
	StashCmd.AddCommand(RestoreCmd)
	StashCmd.AddCommand(DeleteCmd)
	StashCmd.AddCommand(ShowCmd)
}

func Snap(stashMessage string) int {
	exitCode := stashPushAll(stashMessage)
	if exitCode != 0 {
		return exitCode
	}
	return stashApply()
}

func stashPushAll(stashMessage string) int {
	gitSubCmd := []string{"stash", "push", "--include-untracked"}
	if stashMessage != "" {
		commitId := git.GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}
func stashPushOnlyWorktree(stashMessage string) int {
	gitSubCmd := []string{"stash", "push", "--include-untracked", "--keep-index"}
	if stashMessage != "" {
		commitId := git.GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}

func stashPushOnlyIndex(stashMessage string) int {
	gitSubCmd := []string{"stash", "push", "--staged"}
	if stashMessage != "" {
		commitId := git.GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}

func stashApply() int {
	gitSubCmd := []string{"stash", "apply", "--index", "--quiet"}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}
