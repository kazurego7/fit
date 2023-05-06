package branch

import (
	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var RenameCmd = &cobra.Command{
	Use:   "rename <branch name>",
	Short: "現在のブランチの名前を変更する.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), git.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"branch", "--move", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
