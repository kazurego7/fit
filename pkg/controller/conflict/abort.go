package conflict

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var AbortCmd = &cobra.Command{
	Use:   "abort",
	Short: "マージコンフリクトの解消を中止し、ワークツリー・インデックスの変更を元に戻す.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"merge", "--abort"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
