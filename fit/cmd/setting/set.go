package setting

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "設定を追加する.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		scopeFlag := getScopeFlag()
		gitSubCmd := []string{"config", scopeFlag, args[0], args[1]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
		return nil
	},
}
