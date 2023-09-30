package branch

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch <branch>",
	Short: "指定したブランチに移動し、ワークツリー・インデックスを復元する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"switch", args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		branchNameList, err := git.GetBranchNameListInUpdateOrder()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return branchNameList, cobra.ShellCompDirectiveNoFileComp
	},
}
