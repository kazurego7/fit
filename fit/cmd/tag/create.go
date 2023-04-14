package tag

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <new tag name>",
	Short: "現在のコミットにタグを付ける.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"tag", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
