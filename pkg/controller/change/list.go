package change

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "ワークツリー・インデックスの変更があるファイルを一覧表示する.",
	Run: func(cmd *cobra.Command, args []string) {
		var gitSubCmd []string
		switch {
		case listFlag.details:
			gitSubCmd = []string{"--paginate", "-c", "status.relativePaths=false", "status", "--verbose", "--verbose", "--untracked-files=all"}
		default:
			gitSubCmd = []string{"--paginate", "status", "--short", "--untracked-files=all"}
		}
		util.GitCommand(usecase.RootFlag, gitSubCmd...)
	},
}

var listFlag struct {
	details bool
}

func init() {
	ListCmd.Flags().BoolVarP(&listFlag.details, "details", "d", false, "ファイルの変更の詳細を表示する")
}
