package commit

import (
	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/service"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <message>",
	Short: "インデックスから新しいコミットを作成し、現在のブランチをそのコミットに移動する.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		if createFlag.all {
			service.StageChange([]string{":/"})
		}
		git.Commit(args[0])
		git.ShowStatus()
	},
}

var createFlag struct {
	all bool
}

func init() {
	CreateCmd.Flags().BoolVarP(&createFlag.all, "all", "a", false, "ワークツリー・インデックスの両方から新しいコミットを作成する.")
}
