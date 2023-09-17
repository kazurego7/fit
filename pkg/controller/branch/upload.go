package branch

import (
	"fit/pkg/infra/git"
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload [branch name]",
	Short: "ブランチとそのブランチの指すコミットをリモートリポジトリにアップロードする.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var branchName string
		if len(args) == 0 {
			branchName = "HEAD"
		} else {
			branchName = "heads/" + args[0]
		}
		var gitSubCmd []string
		if !git.ExistsUpstreamFor(branchName) {
			// すでに upstream が設定されている場合は、upstream を設定しない
			gitSubCmd = []string{"push", "origin", branchName, "--prune", "--set-upstream"}
		} else {
			gitSubCmd = []string{"push", "origin", branchName, "--prune"}
		}
		util.GitCommand(usecase.RootFlag, gitSubCmd)
	},
}
