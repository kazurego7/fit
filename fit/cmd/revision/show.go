package revision

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "指定したコミットに含まれるファイルを表示する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"show", "--stat", "--summary", "--patch", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
