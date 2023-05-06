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
	Use:   "stage <pathspec>…",
	Short: "ワークツリーのファイルの変更をインデックスにステージングする.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), git.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		// コンフリクト解消していないファイルがあればステージングしない
		if err := checkConflictResolved(args...); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		// index にも worktree にもあるファイルは上書き対象となる
		indexList := git.SearchIndexList("u", args...)
		overwriteList := git.SearchWorktreeList("", indexList...)

		// index への上書きがある場合は、バックアップを行う
		if len(overwriteList) != 0 {
			git.Snap(`"fit change stage" のバックアップ`, args...)
			fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
				`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
		}
		gitSubCmd := append([]string{"add"}, args...)
		util.GitCommand(global.RootFlag, gitSubCmd...)
	},
	ValidArgs: append(git.SearchUntrackedFiles(":/"), git.SearchWorktreeList("u", ":/")...),
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
