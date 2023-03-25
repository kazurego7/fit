package branch

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/spf13/cobra"
)

var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// すでに upstream が設定されている場合は、upstream を設定しない
		var gitSubCmd []string
		if !existsUpstreamBranch(pushFlags.branch) {
			gitSubCmd = []string{"push", "origin", pushFlags.branch, "--prune", "--set-upstream"}
		} else {
			gitSubCmd = []string{"push", "origin", pushFlags.branch, "--prune"}
		}
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
	},
}

var pushFlags struct {
	branch string
}

func init() {
	PushCmd.Flags().StringVarP(&pushFlags.branch, "branch", "b", "HEAD", "choose branch name or HEAD")
}

func existsUpstreamBranch(branchName string) bool {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", " --symbolic-full-name", `"` + branchName + `@{u}"`}
	_, err := fitio.ExecuteGitOutput(gitSubCmd...)
	return err != nil
}
