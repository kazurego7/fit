package branch

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
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
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		flagBranch := getBranchName(pullFlags.branch)
		exitCode := pullFor(flagBranch)
		if exitCode != 0 {
			return
		}
		if !existsUpstreamFor(flagBranch) {
			setUpstreamTo(flagBranch)
		}
	},
}

var pullFlags struct {
	branch string
}

func init() {
	PullCmd.Flags().StringVarP(&pullFlags.branch, "branch", "b", "HEAD", "choose branch name or HEAD")
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
	fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}

func setUpstreamTo(branch string) int {
	gitSubCmd := []string{"branch", branch, "--set-upstream-to=origin/" + branch}
	fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
	exitCode := fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	return exitCode
}
