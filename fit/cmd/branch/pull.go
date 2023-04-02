package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
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
		gitSubCmd = []string{"pull", "origin", currentBranch, "--ff-only", "--prune"}
	} else {
		gitSubCmd = []string{"fetch", "origin", branch + ":" + branch, "--ff-only", "--prune"}
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
