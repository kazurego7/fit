package repository

import (
	"fit/pkg/global"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var FetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "リモートリポジトリからブランチ・タグ・コミットを取得する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"fetch", "origin", "--prune"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}