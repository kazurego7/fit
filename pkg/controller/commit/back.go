package commit

import (
	"github.com/spf13/cobra"
)

var BackCmd = &cobra.Command{
	Use:   "back",
	Short: "ワークツリー・インデックスを復元せずに、1つ前のコミットに移動する.",
	Args:  cobra.MatchAll(cobra.NoArgs, service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		git.ResetHeadWithoutWorktreeAndIndex()
		git.ShowStatus()
	},
}
