package change

import (
	"errors"

	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UnstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "インデックスにステージングされているファイルの変更をワークツリーに戻す.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// index にも worktree にもあるファイルは上書き対象となる
		indexList := git.SearchIndexList("", args[0])
		overwriteList := git.SearchWorktreeList("", indexList...)

		// worktree への上書きがある場合は、バックアップを促す
		if len(overwriteList) != 0 {
			confirmBackup()
			exitCode := restoreWorktree(overwriteList...)
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
		gitSubCmd := []string{"diff", "--cached", "--name-only", "--relative"}
		out, _, err := util.GitQuery(global.RootFlag, gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := util.SplitLn(string(out))
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
