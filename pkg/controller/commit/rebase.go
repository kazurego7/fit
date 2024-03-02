package commit

import (
	"github.com/kazurego7/fit/pkg/domain"
	"github.com/spf13/cobra"
)

var RebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "現在のブランチのコミットをメインラインに移動する.",
	Args:  cobra.MatchAll(cobra.NoArgs, service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		mainline := git.GetFitConfig(domain.FitConfigConstant.MainlineType())
		git.RebaseToMainline(mainline)
	},
}
