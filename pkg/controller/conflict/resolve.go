package conflict

import (
	"fit/pkg/global"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var ResolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "マージコンフリクトを解消し、マージコミットを作成する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"merge", "--continue"}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}
