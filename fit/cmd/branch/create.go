package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "現在のコミットにブランチを作成し、そのブランチに移動する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"switch", "--create", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
