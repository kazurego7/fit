package history

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "指定したブランチに移動する.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"switch", args[0]}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"for-each-ref", `--format="%(refname:short)"`, "refs/remotes", "refs/heads"}
		out, _, err := util.GitQuery(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := util.SplitLn(string(out))
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
