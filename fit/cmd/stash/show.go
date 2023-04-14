package stash

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show <stash>",
	Short: "指定したスタッシュに格納されているファイルの内容を表示する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"stash", "show", "--stat", "--summary", "--patch", "--include-untracked", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
