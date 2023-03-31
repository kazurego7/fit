package branch

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
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		var gitSubCmd []string
		if !existsUpstreamFor(pushFlags.branch) {
			// すでに upstream が設定されている場合は、upstream を設定しない
			gitSubCmd = []string{"push", "origin", pushFlags.branch, "--prune", "--set-upstream"}
		} else {
			gitSubCmd = []string{"push", "origin", pushFlags.branch, "--prune"}
		}
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}

var pushFlags struct {
	branch string
}

func init() {
	PushCmd.Flags().StringVarP(&pushFlags.branch, "branch", "b", "HEAD", "choose branch name or HEAD")
}
