package change

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var StageCmd = &cobra.Command{
	Use:   "stage",
	Short: "ワークツリーのファイルの変更をインデックスにステージングする.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// index にも worktree にもあるファイルは上書き対象となる
		indexList := searchIndexList("", args[0])
		overwriteList := searchWorktreeList("", indexList...)

		// index への上書きがある場合は、バックアップを促す
		if len(overwriteList) != 0 {
			confirmBackup()
		}
		gitSubCmd := []string{"add", args[0]}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"ls-files", `--modified`, "--others"}
		out, _, err := util.GitQuery(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := util.SplitLn(string(out))
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
