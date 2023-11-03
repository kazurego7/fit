package branch

import (
	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/service"

	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch <branch>",
	Short: "指定したブランチに移動し、ワークツリー・インデックスを復元する(作業中のファイルは一時保存する).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		service.SwitchBranchAfterWIP(args[0])
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		branchNameList, err := git.GetBranchNameListInUpdateOrder()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return branchNameList, cobra.ShellCompDirectiveNoFileComp
	},
}
