package change

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var StageCmd = &cobra.Command{
	Use:   "stage",
	Short: "ワークツリーのファイルの変更をインデックスにステージングする.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// コンフリクト解消していないファイルがあればステージングしない
		if err := checkConflictResolved(args...); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		// index にも worktree にもあるファイルは上書き対象となる
		indexList := git.SearchIndexList("", args...)
		overwriteList := git.SearchWorktreeList("", indexList...)

		// index への上書きがある場合は、バックアップを促す
		if len(overwriteList) != 0 {
			confirmBackup()
		}
		gitSubCmd := append([]string{"add"}, args...)
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"ls-files", `--modified`, "--others"}
		out, _, err := util.GitQuery(global.RootFlag, gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := util.SplitLn(string(out))
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}

func checkConflictResolved(args ...string) error {
	gitSubCmd := append(
		[]string{
			"-c",
			"core.whitespace=-trailing-space,-space-before-tab,-indent-with-non-tab,-tab-in-indent,-cr-at-eol",
			"diff",
			"--check",
		},
		args...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	if string(out) != "" {
		unmergedList := git.SearchWorktreeList("U", args...)
		errorMessage := "コンフリクトマーカーが残っています. コンフリクトマーカーを取り除いてください\n" + strings.Join(unmergedList, "\n")
		return errors.New(errorMessage)
	}
	return nil
}
