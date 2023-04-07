package history

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "インデックスから新しいgitリビジョンを作成し、現在のブランチをそのリビジョンに移動する.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := append([]string{"commit", "--message"}, args...)
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}
