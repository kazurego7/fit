package branch

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <new branch name>",
	Short: "指定したコミットにブランチを作成し、そのブランチに移動する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"checkout", "-b", args[0], createFlag.revision}
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
}

var createFlag struct {
	revision string
}

func init() {
	CreateCmd.Flags().StringVarP(&createFlag.revision, "revision", "r", "HEAD", "ブランチ作成先のコミット")
}
