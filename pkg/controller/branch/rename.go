package branch

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/service"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var RenameCmd = &cobra.Command{
	Use:   "rename <branch name>",
	Short: "現在のブランチの名前を変更する.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"branch", "--move", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
