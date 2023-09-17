package change

import (
	"fmt"
	"os"

	"fit/pkg/global"
	"fit/pkg/infra/git"
	"fit/pkg/service"
	"fit/pkg/util"

	"github.com/spf13/cobra"
)

var StageCmd = &cobra.Command{
	Use:   "stage <pathspec>…",
	Short: "ワークツリーのファイルの変更をインデックスにステージングする.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		// コンフリクト解消していないファイルがあればステージングしない
		if err := service.CheckConflictResolved(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		// index にも worktree にもあるファイルは上書き対象となる
		indexList := git.SearchIndexList("u", args)
		overwriteList := git.SearchWorktreeList("", indexList)

		// index への上書きがある場合は、バックアップを行う
		if len(overwriteList) != 0 {
			service.Snap(`"fit change stage" のバックアップ`, args...)
			fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
				`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
		}
		gitSubCmd := append([]string{"add"}, args...)
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
	ValidArgs: append(git.SearchUntrackedFiles([]string{":/"}), git.SearchWorktreeList("u", []string{":/"})...),
}
