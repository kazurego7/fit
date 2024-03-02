package branch

import (
	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch <branch>",
	Short: "指定したブランチに移動し、ワークツリー・インデックスを復元する(作業中のファイルは一時保存する).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		if switchFlag.noWip {
			git.SwitchBranch(branch)
		} else {
			service.SwitchBranchAfterWIP(branch)
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		branchNameList, err := git.GetBranchNameListInUpdateOrder()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return branchNameList, cobra.ShellCompDirectiveNoFileComp
	},
}

var switchFlag struct {
	noWip bool
}

func init() {
	SwitchCmd.Flags().BoolVar(&switchFlag.noWip, "no-wip", false, "作業中のファイルを一時保存・復元せずにブランチを移動する")
}
