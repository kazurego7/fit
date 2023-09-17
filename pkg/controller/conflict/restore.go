package conflict

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <pathspec>",
	Short: "ファイルをマージコンフリクトの状態に戻す.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"restore", "--merge", args[0]}
		util.GitCommand(usecase.RootFlag, gitSubCmd...)
	},
}
