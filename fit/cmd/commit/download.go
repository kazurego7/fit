package commit

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "リモートリポジトリからブランチ・タグ・コミットをダウンロードする.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"fetch", "origin", "--prune"}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
