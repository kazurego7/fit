package conflict

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "マージコンフリクトのあるファイルを一覧表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"diff", "--name-only", "--diff-filter=U"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
