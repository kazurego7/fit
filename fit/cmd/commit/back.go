package commit

import (
	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var BackCmd = &cobra.Command{
	Use:   "back",
	Short: "現在のブランチを1つ前のコミットに移動する.",
	Args:  cobra.MatchAll(cobra.NoArgs, git.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"reset", "--soft", "HEAD^"}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
