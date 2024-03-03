package branch

import (
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
		// upstream が設定されていない場合は、upstream を設定
		if !git.ExistsUpstreamFor(branchName) {
			git.SetUpstream(branchName)
			git.PushFor(branchName)
		} else {
			git.PushFor(branchName)
		}
	},
}
