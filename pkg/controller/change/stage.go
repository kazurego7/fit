package change

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"fit/pkg/infra"
	"fit/pkg/usecase"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var StageCmd = &cobra.Command{
	Use:   "stage <pathspec>…",
	Short: "ワークツリーのファイルの変更をインデックスにステージングする.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), infra.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		// コンフリクト解消していないファイルがあればステージングしない
		if err := checkConflictResolved(args...); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		// index にも worktree にもあるファイルは上書き対象となる
		indexList := infra.SearchIndexList("u", args...)
		overwriteList := infra.SearchWorktreeList("", indexList...)

		// index への上書きがある場合は、バックアップを行う
		if len(overwriteList) != 0 {
			infra.Snap(`"fit change stage" のバックアップ`, args...)
			fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
				`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
		}
		gitSubCmd := append([]string{"add"}, args...)
		util.GitCommand(usecase.RootFlag, gitSubCmd...)
	},
	ValidArgs: append(infra.SearchUntrackedFiles(":/"), infra.SearchWorktreeList("u", ":/")...),
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
	out, _, _ := util.GitQuery(usecase.RootFlag, gitSubCmd...)
	if string(out) != "" {
		unmergedList := infra.SearchWorktreeList("U", args...)
		errorMessage := "コンフリクトマーカーが残っています. コンフリクトマーカーを取り除いてください\n" + strings.Join(unmergedList, "\n")
		return errors.New(errorMessage)
	}
	return nil
}
