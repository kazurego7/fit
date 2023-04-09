package conflict

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var CompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "マージコンフリクトの解消を完了し、マージコミットを作成する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"merge", "--continue"}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}
