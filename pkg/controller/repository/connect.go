package repository

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect <remote url>",
	Short: "リモートリポジトリとの非同期接続を設定する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"remote", "add", "origin", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
