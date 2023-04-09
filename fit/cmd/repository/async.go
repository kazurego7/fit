package repository

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var AsyncCmd = &cobra.Command{
	Use:   "async",
	Short: "リモートリポジトリとの非同期接続を設定する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"remote", "add", "origin", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
