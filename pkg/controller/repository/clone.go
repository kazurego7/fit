package repository

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var CloneCmd = &cobra.Command{
	Use:   "clone <remote url>",
	Short: "リモートリポジトリのクローンをローカルに作成する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"clone", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
