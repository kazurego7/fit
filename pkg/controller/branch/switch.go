package branch

import (
	"strings"

	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch <branch>",
	Short: "指定したブランチに移動し、ワークツリー・インデックスを復元する(作業中のファイルは一時保存する).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		const WIP_MESSAGE = "[WIP]"

		if git.ExistsIndexDiff([]string{":/"}) {
			git.Commit(WIP_MESSAGE + " Index")
		}

		if git.ExistsUntrackedFiles([]string{":/"}) || git.ExistsWorktreeDiff([]string{":/"}) {
			gitSubCmd := []string{"add", ":/"}
			util.GitCommand(global.RootFlag, gitSubCmd)
			git.Commit(WIP_MESSAGE + " Worktree")
		}

		{
			gitSubCmd := []string{"switch", args[0]}
			util.GitCommand(global.RootFlag, gitSubCmd)
		}

		var existsWIPIndex = false
		{
			gitSubCmd := []string{"log", "--format=%B -n 1", "HEAD"}
			out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
			existsWIPIndex = strings.HasPrefix(string(out), WIP_MESSAGE)
		}
		if existsWIPIndex {
			gitSubCmd := []string{"reset", "--mixed", "HEAD^"}
			util.GitCommand(global.RootFlag, gitSubCmd)
		}

		var existsWIPWorktree = false
		{
			gitSubCmd := []string{"log", "--format=%B -n 1", "HEAD"}
			out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
			existsWIPWorktree = strings.HasPrefix(string(out), WIP_MESSAGE)
		}
		if existsWIPWorktree {
			gitSubCmd := []string{"reset", "--soft", "HEAD^"}
			util.GitCommand(global.RootFlag, gitSubCmd)
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		branchNameList, err := git.GetBranchNameListInUpdateOrder()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return branchNameList, cobra.ShellCompDirectiveNoFileComp
	},
}
