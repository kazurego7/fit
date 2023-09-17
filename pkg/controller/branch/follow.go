package branch

import (
	"errors"

	"fit/pkg/infra"
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var FollowCmd = &cobra.Command{
	Use:   "follow [<branch name>]",
	Short: "ローカルブランチをリモートブランチの状態に追従させる.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 && infra.ShowCurrentBranch() == "" {
			return errors.New("現在のブランチが選択されていません.\n" +
				"※ ブランチを指定するか、\"fit branch switch\" でブランチの切り替えをしてください")
		}

		var branchName string
		if len(args) == 0 {
			branchName = "HEAD"
		} else {
			branchName = args[0]
		}

		flagBranch := getBranchName(branchName)
		pullForExitCode := pullFor(flagBranch)
		if pullForExitCode != 0 {
			return errors.New("ブランチの取得に失敗しました")
		}
		if !existsUpstreamFor(flagBranch) {
			setUpstreamExitCode := setUpstream(flagBranch)
			if setUpstreamExitCode != 0 {
				return errors.New("リモートリポジトリのブランチの設定に失敗しました")
			}
		}
		switchBranch(branchName)
		return nil
	},
}

func pullFor(branch string) int {
	// pull したいブランチにチェックアウトしているか、そうでないかで処理を分岐
	currentBranch := getBranchName("HEAD")
	var gitSubCmd []string
	if currentBranch == branch {
		gitSubCmd = []string{"pull", "origin", currentBranch + ":" + currentBranch, "--ff-only", "--prune"}
	} else {
		gitSubCmd = []string{"fetch", "origin", branch + ":" + branch, "--prune"}
	}
	exitCode := util.GitCommand(usecase.RootFlag, gitSubCmd...)
	return exitCode
}

func setUpstream(branch string) int {
	gitSubCmd := []string{"branch", branch, "--set-upstream-to=origin/" + branch}
	exitCode := util.GitCommand(usecase.RootFlag, gitSubCmd...)
	return exitCode
}

func switchBranch(branch string) int {
	gitSubCmd := []string{"switch", branch}
	exitCode := util.GitCommand(usecase.RootFlag, gitSubCmd...)
	return exitCode
}
