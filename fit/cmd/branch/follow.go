package branch

import (
	"errors"

	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
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

		flagBranch := getBranchName(branchName)
		exitCode := pullFor(flagBranch)
		if exitCode != 0 {
			return errors.New("ブランチの取得に失敗しました")
		}
		if !existsUpstreamFor(flagBranch) {
			setUpstreamTo(flagBranch)
		}
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
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}

func setUpstreamTo(branch string) int {
	gitSubCmd := []string{"branch", branch, "--set-upstream-to=origin/" + branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}
