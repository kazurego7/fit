package change

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Diff files with changes in the index or worktree",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var gitSubCmd []string
		if showFlags.index {
			gitSubCmd = []string{"diff", "--staged", args[0]}
		} else {
			gitSubCmd = []string{"diff", args[0]}
		}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var gitSubCmd []string
		if showFlags.index {
			gitSubCmd = []string{"diff", "--staged", "--name-only", "--relative"}
		} else {
			gitSubCmd = []string{"diff", "--name-only", "--relative"}
		}
		out, _, err := util.GitQuery(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := util.SplitLn(string(out))
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}

var showFlags struct {
	worktree bool
	index    bool
}

func init() {
	ShowCmd.Flags().BoolVarP(&showFlags.worktree, "worktree", "w", false, "diff worktree to index")
	ShowCmd.Flags().BoolVarP(&showFlags.index, "index", "i", false, "diff index to HEAD")
	ShowCmd.MarkFlagsMutuallyExclusive("worktree", "index")
}
