package tag

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload <tag>",
	Short: "リモートリポジトリにタグをアップロードする(警告：アップロードされたタグは削除できません).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"push", "origin", "tags/" + args[0], "--prune"}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
