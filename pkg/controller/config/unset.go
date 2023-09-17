package config

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var UnsetCmd = &cobra.Command{
	Use:   "unset <name>",
	Short: "設定を削除する.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scopeFlag := getScopeFlag()
		gitSubCmd := []string{"config", scopeFlag, "--unset", args[0]}
		util.GitCommand(usecase.RootFlag, gitSubCmd)
		return nil
	},
}
