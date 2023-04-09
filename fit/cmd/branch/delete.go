package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "指定したブランチを削除する(ブランチの指すコミットは削除しない).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"branch", "--delete", "--force", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
