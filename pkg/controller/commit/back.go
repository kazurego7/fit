package commit

import (
	service "github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra/git"

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
