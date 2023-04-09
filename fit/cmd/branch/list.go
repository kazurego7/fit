package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "ブランチとその上流ブランチを一覧表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"branch", "--list", "--verbose", "--verbose"}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
