package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload [branch name]",
	Short: "ブランチとそのブランチの指すリビジョンをリモートリポジトリにアップロードする.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var branchName string
		if len(args) == 0 {
			branchName = "HEAD"
		} else {
			branchName = "heads/" + args[0]
		}
		var gitSubCmd []string
		if !existsUpstreamFor(branchName) {
			// すでに upstream が設定されている場合は、upstream を設定しない
			gitSubCmd = []string{"push", "origin", branchName, "--prune", "--set-upstream"}
		} else {
			gitSubCmd = []string{"push", "origin", branchName, "--prune"}
		}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
