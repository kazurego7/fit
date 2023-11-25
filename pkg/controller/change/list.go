package change

import (
	"github.com/kazurego7/fit/pkg/infra/git"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "ワークツリー・インデックスの変更があるファイルを一覧表示する.",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case listFlag.details:
			git.ShowChangeDetails()
		default:
			git.ShowStatus()
		}
	},
}

var listFlag struct {
	details bool
}

func init() {
	ListCmd.Flags().BoolVarP(&listFlag.details, "details", "d", false, "ファイルの変更の詳細を表示する")
}
