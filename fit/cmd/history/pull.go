package history

import (
	"strings"

	"github.com/kazurego7/fit/fit/fitio"
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
	Run: func(cmd *cobra.Command, args []string) {
		currentBranch := getBranchName("HEAD")
		flagBranch := getBranchName(pullFlags.branch)

		// pull したいブランチにチェックアウトしているか、そうでないかで処理を分岐
		var gitSubCmd []string
		if currentBranch == flagBranch {
			gitSubCmd = []string{"pull", "origin", currentBranch, "--prune"}
		} else {
			gitSubCmd = []string{"fetch", "origin", flagBranch + ":" + flagBranch, "--prune"}
		}
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
	},
}

var pullFlags struct {
	branch string
}

func init() {
	PullCmd.Flags().StringVarP(&pullFlags.branch, "branch", "b", "HEAD", "choose branch name or HEAD")
}

func getBranchName(refspec string) string {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", refspec}
	out, _ := fitio.ExecuteGitOutput(gitSubCmd...)
	return strings.Trim(string(out), "\n")

}
