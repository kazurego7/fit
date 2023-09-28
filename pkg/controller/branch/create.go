package branch

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <new branch name>",
	Short: "指定したコミットにブランチを作成し、そのブランチに移動する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"checkout", "-b", args[0], createFlag.commit}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}

var createFlag struct {
	commit string
}

func init() {
	CreateCmd.Flags().StringVarP(&createFlag.commit, "commit", "c", "HEAD", "ブランチ作成先のコミット")
}
