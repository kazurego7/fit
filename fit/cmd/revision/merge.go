package revision

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var MergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "指定したブランチを現在のブランチにマージする.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"merge", args[0]}
		allArgs := append(gitSubCmd, args...)
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, allArgs...)
	},
}
