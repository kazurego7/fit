package repository

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var RemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "リモートリポジトリを一覧表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"remote", "--verbose"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
