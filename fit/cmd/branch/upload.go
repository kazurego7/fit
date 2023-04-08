package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "ブランチとそのブランチの指すコミットをリモートリポジトリにアップロードする.",
	Args:  cobra.MaximumNArgs(1),
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
			gitSubCmd = []string{"push", "origin", "heads/" + branchName, "--prune", "--set-upstream"}
		} else {
			gitSubCmd = []string{"push", "origin", "heads/" + branchName, "--prune"}
		}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}
