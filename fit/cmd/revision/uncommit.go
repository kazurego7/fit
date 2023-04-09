package revision

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UncommitCmd = &cobra.Command{
	Use:   "uncommit",
	Short: "現在のブランチを1つ前のコミットに移動する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"reset", "--soft", "HEAD^"}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
