package tag

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload <tag>",
	Short: "リモートリポジトリにタグをアップロードする(警告：アップロードされたタグは削除できません).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"push", "origin", "tags/" + args[0], "--prune"}
		util.GitCommand(usecase.RootFlag, gitSubCmd)
	},
}
