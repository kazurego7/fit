package commit

import (
	service "github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra/git"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <message>",
	Short: "インデックスから新しいコミットを作成し、現在のブランチをそのコミットに移動する.",
	Args:  service.CurrentIsNotReadonly(),
	Run: func(cmd *cobra.Command, args []string) {
		if createFlag.all {
			service.StageChange([]string{":/"})
		}

		if len(args) == 0 {
			git.CommitWithOpenEditor()
		} else {
			git.CommitWithNoAllowEmpty(args[0])
		}
		git.ShowStatus()
	},
}

var createFlag struct {
	all bool
}

func init() {
	CreateCmd.Flags().BoolVarP(&createFlag.all, "all", "a", false, "ワークツリー・インデックスの両方から新しいコミットを作成する.")
}
