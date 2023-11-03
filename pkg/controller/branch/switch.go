package branch

import (
	"strings"

	"github.com/kazurego7/fit/pkg/infra/git"

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
			git.AddStageing([]string{":/"})
			git.Commit(WIP_MESSAGE + " Worktree")
		}

		git.SwitchBranch(args[0])

		if strings.HasPrefix(git.GetCommitMessage("HEAD"), WIP_MESSAGE) {
			git.ResetHeadWithoutWorktree()
		}
		if strings.HasPrefix(git.GetCommitMessage("HEAD"), WIP_MESSAGE) {
			git.ResetHeadWithoutWorktreeAndIndex()
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
