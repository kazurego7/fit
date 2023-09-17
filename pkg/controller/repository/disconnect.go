package repository

import (
	"fit/pkg/global"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var DisconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "リモートリポジトリとの非同期接続をやめる.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"remote", "remove", "origin"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
