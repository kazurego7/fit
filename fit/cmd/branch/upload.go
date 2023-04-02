package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload",
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
		var gitSubCmd []string
		if !existsUpstreamFor(branchName) {
			// すでに upstream が設定されている場合は、upstream を設定しない
			gitSubCmd = []string{"push", "origin", branchName, "--prune", "--set-upstream"}
		} else {
			gitSubCmd = []string{"push", "origin", branchName, "--prune"}
		}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}
