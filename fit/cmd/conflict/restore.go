package conflict

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <pathspec>",
	Short: "ファイルをマージコンフリクトの状態に戻す.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"restore", "--merge", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
