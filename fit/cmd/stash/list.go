package stash

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "保存されたスタッシュを一覧表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"reflog", "show", "--format=%C(03)%h%C(reset) %C(bold 1)%gD%C(reset) %C(bold 0)%s%C(reset)", "stash"}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
