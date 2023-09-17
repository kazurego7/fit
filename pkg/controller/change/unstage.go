package change

import (
	"errors"
	"fmt"

	"fit/pkg/infra/git"
	"fit/pkg/service"

	"github.com/spf13/cobra"
)

var UnstageCmd = &cobra.Command{
	Use:   "unstage <pathspec>…",
	Short: "インデックスにステージングされているファイルの変更をワークツリーに戻す.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	RunE: func(cmd *cobra.Command, args []string) error {
		// index にも worktree にもあるファイルは上書き対象となる
		indexList := git.SearchIndexList("", args)
		overwriteList := git.SearchWorktreeList("", indexList)

		// worktree への上書きがある場合は、バックアップを行う
		if len(overwriteList) != 0 {
			service.Snap(`"fit change unstage" のバックアップ`, args...)
			fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
				`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
			exitCode := git.RestoreWorktree(overwriteList)
			if exitCode != 0 {
				return errors.New("restore index failed")
			}
			git.RestoreIndex(args)
		} else {
			git.RestoreIndex(args)
		}
		return nil
	},
	ValidArgs: git.SearchIndexList("", []string{":/"}),
}
