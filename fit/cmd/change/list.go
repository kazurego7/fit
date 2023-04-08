package change

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "ワークツリー・インデックスの変更があるファイルを一覧表示する.",
	Run: func(cmd *cobra.Command, args []string) {
		var gitSubCmd []string
		switch {
		case listFlag.details:
			gitSubCmd = []string{"-p", "-c", "status.relativePaths=false", "status", "--verbose", "--verbose", "--untracked-files=all"}
		case listFlag.all:
			gitSubCmd = []string{"status", "--short", "--untracked-files=all"}
		default:
			gitSubCmd = []string{"status", "--short"}
		}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}

var listFlag struct {
	all     bool
	details bool
}

func init() {
	ListCmd.Flags().BoolVarP(&listFlag.all, "all", "a", false, "追跡されていないディレクトリにある個々のファイルも表示する")
	ListCmd.Flags().BoolVarP(&listFlag.details, "details", "d", false, "ファイルの変更の詳細を表示する")
	ListCmd.MarkFlagsMutuallyExclusive("all", "details")
}
