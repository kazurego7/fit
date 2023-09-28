package commit

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch <commit>",
	Short: "指定したコミットに読み取り専用で切り替える.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := append([]string{"switch", "--detach"}, args[0])
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
