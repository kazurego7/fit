package branch

import (
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch <branch>",
	Short: "指定したブランチに移動する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"switch", args[0]}
		util.GitCommand(usecase.RootFlag, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"for-each-ref", `--format="%(refname:short)"`, "refs/remotes", "refs/heads"}
		out, _, err := util.GitQuery(usecase.RootFlag, gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := util.SplitLn(string(out))
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
