package branch

import (
	"errors"

	"fit/pkg/infra/git"
	"fit/pkg/service"

	"github.com/spf13/cobra"
)

var FollowCmd = &cobra.Command{
	Use:   "follow [<branch name>]",
	Short: "ローカルブランチをリモートブランチの状態に追従させる.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 && git.ShowCurrentBranch() == "" {
			return errors.New("現在のブランチが選択されていません.\n" +
				"※ ブランチを指定するか、\"fit branch switch\" でブランチの切り替えをしてください")
		}

		var branchName string
		if len(args) == 0 {
			branchName = "HEAD"
		} else {
			branchName = args[0]
		}

		flagBranch := git.GetBranchName(branchName)
		pullForExitCode := git.PullFor(flagBranch)
		if pullForExitCode != 0 {
			return errors.New("ブランチの取得に失敗しました")
		}
		if !git.ExistsUpstreamFor(flagBranch) {
			setUpstreamExitCode := git.SetUpstream(flagBranch)
			if setUpstreamExitCode != 0 {
				return errors.New("リモートリポジトリのブランチの設定に失敗しました")
			}
		}
		git.SwitchBranch(branchName)
		git.FetchPrune()
		service.PruneBranchOfGone()
		return nil
	},
}
