package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var FollowCmd = &cobra.Command{
	Use:   "follow",
	Short: "ローカルブランチをリモートブランチの状態に追従させる.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var branchName string
		if len(args) == 0 {
			branchName = "HEAD"
		} else {
			branchName = args[0]
		}

		flagBranch := getBranchName(branchName)
		exitCode := pullFor(flagBranch)
		if exitCode != 0 {
			return
		}
		if !existsUpstreamFor(flagBranch) {
			setUpstreamTo(flagBranch)
		}
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
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}

func setUpstreamTo(branch string) int {
	gitSubCmd := []string{"branch", branch, "--set-upstream-to=origin/" + branch}
	util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}
