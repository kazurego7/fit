package change

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "ワークツリーおよびインデックスの変更があるファイルを一覧表示する.",
	Run: func(cmd *cobra.Command, args []string) {
		var gitSubCmd []string
		if listFlag.all {
			gitSubCmd = []string{"status", "--short", "--untracked-files=all"}
		} else {
			gitSubCmd = []string{"status", "--short"}
		}

		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}

var listFlag struct {
	all bool
}

func init() {
	ListCmd.Flags().BoolVarP(&listFlag.all, "all", "a", false, "追跡されていないディレクトリにある個々のファイルも表示する")
}
