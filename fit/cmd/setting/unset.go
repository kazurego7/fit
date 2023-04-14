package setting

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UnsetCmd = &cobra.Command{
	Use:   "unset <name>",
	Short: "設定を削除する.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scopeFlag := getScopeFlag()
		gitSubCmd := []string{"config", scopeFlag, "--unset", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
		return nil
	},
}
