package change

import (
	"errors"

	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UnstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var restoreList []string
		indexList := searchIndexList("", args[0])
		if indexList[0] != "" {
			restoreList = searchWorktreeList("", indexList...)
		}
		if restoreList[0] != "" {
			exitCode := restoreWorktree(restoreList...)
			if exitCode != 0 {
				return errors.New("restore index failed")
			}
			restoreIndex(args[0])
		} else {
			restoreIndex(args[0])
		}
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"diff", "--cached", "--name-only"}
		out, _, err := util.GitQuery(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := util.SplitLn(string(out))
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
