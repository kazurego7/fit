package change

import (
	"errors"
	"fmt"

	"github.com/kazurego7/fit/fit/git"
	"github.com/spf13/cobra"
)

var UnstageCmd = &cobra.Command{
	Use:   "unstage <pathspec>…",
	Short: "インデックスにステージングされているファイルの変更をワークツリーに戻す.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), git.CurrentIsNotReadonly()),
	RunE: func(cmd *cobra.Command, args []string) error {
		// index にも worktree にもあるファイルは上書き対象となる
		indexList := git.SearchIndexList("", args...)
		overwriteList := git.SearchWorktreeList("", indexList...)

		// worktree への上書きがある場合は、バックアップを行う
		if len(overwriteList) != 0 {
			git.Snap(`"fit change unstage" のバックアップ`, args...)
			fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
				`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
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
	ValidArgs: git.SearchIndexList("", ":/"),
}
