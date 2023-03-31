package index

import (
	"strings"

	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var DiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var gitSubCmd []string
		if diffFlags.indexToHead {
			gitSubCmd = []string{"diff", "--staged", args[0]}
		} else {
			gitSubCmd = []string{"diff", args[0]}
		}
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var gitSubCmd []string
		if diffFlags.indexToHead {
			gitSubCmd = []string{"diff", "--staged", "--name-only", "--relative"}
		} else {
			gitSubCmd = []string{"diff", "--name-only", "--relative"}
		}
		out, _, err := fitio.GitQuery(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := strings.Split(strings.Trim(strings.ReplaceAll(string(out), `"`, ""), "\n"), "\n")
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}

var diffFlags struct {
	worktreeToIndex bool
	indexToHead     bool
}

func init() {
	DiffCmd.Flags().BoolVarP(&diffFlags.worktreeToIndex, "worktree-to-index", "w", false, "diff worktree to index")
	DiffCmd.Flags().BoolVarP(&diffFlags.indexToHead, "index-to-head", "i", false, "diff index to HEAD")
	DiffCmd.MarkFlagsMutuallyExclusive("worktree-to-index", "index-to-head")
}
