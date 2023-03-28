package history

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
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

		var gitSubCmd []string
		if pushFlags.tag != "" {
			// タグ名が設定されていた場合
			gitSubCmd = []string{"push", "origin", "--tags", pushFlags.tag, "--prune"}
		} else {
			// ブランチ名が設定されていた場合
			if !existsUpstreamBranch(pushFlags.branch) {
				// すでに upstream が設定されている場合は、upstream を設定しない
				gitSubCmd = []string{"push", "origin", pushFlags.branch, "--prune", "--set-upstream"}
			} else {
				gitSubCmd = []string{"push", "origin", pushFlags.branch, "--prune"}
			}
		}

		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.CommandGit(global.Flags.Dryrun, gitSubCmd...)
	},
}

var pushFlags struct {
	branch string
	tag    string
}

func init() {
	PushCmd.Flags().StringVarP(&pushFlags.branch, "branch", "b", "HEAD", "choose branch name or HEAD")
	PushCmd.Flags().StringVarP(&pushFlags.tag, "tag", "t", "", "choose tag name")
	PushCmd.MarkFlagsMutuallyExclusive("branch", "tag")
}

func existsUpstreamBranch(branchName string) bool {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", " --symbolic-full-name", `"` + branchName + `@{u}"`}
	_, err := fitio.QueryGit(gitSubCmd...)
	return err != nil
}
