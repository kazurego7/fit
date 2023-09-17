package config

import (
	"fit/pkg/global"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "設定ファイルをエディターで開く.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		scopeFlag := getScopeFlag()
		gitSubCmd := []string{"config", scopeFlag, "--edit"}
		util.GitCommand(global.RootFlag, gitSubCmd)
		return nil
	},
}
