package tag

import (
	"fit/pkg/global"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <tag>",
	Short: "指定されたタグを削除する(タグの指すコミットは削除しない).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"tag", "--delete", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
