package commit

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/service"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <message>",
	Short: "インデックスから新しいコミットを作成し、現在のブランチをそのコミットに移動する.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := append([]string{"commit", "--message"}, args...)
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
