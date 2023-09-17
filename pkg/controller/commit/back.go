package commit

import (
	"fit/pkg/global"
	"fit/pkg/service"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var BackCmd = &cobra.Command{
	Use:   "back",
	Short: "現在のブランチを1つ前のコミットに移動する.",
	Args:  cobra.MatchAll(cobra.NoArgs, service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"reset", "--soft", "HEAD^"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
