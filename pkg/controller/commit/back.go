package commit

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/service"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var BackCmd = &cobra.Command{
	Use:   "back",
	Short: "ワークツリー・インデックスを復元せずに、1つ前のコミットに移動する.",
	Args:  cobra.MatchAll(cobra.NoArgs, service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"reset", "--soft", "HEAD^"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
