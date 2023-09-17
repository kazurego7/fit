package config

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "現在設定されている項目を一覧表示する.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		scopeFlag := getScopeFlag()
		gitSubCmd := []string{"config", "--list", scopeFlag}
		util.GitCommand(usecase.RootFlag, gitSubCmd)
		return nil
	},
}
