package change

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var UnstageCmd = &cobra.Command{
	Use:   "unstage <filename>…",
	Short: "インデックスにステージングされているファイルの変更をワークツリーに戻す.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), service.CurrentIsNotReadonly()),
	RunE: func(cmd *cobra.Command, args []string) error {
		// ファイル名からあいまい検索のパスを作成
		pathList := service.AddFuzzyParentPath(args)
		if unstageFlag.details {
			git.DiffIndex(pathList)
			return nil
		} else {
			// index にも worktree にもあるファイルは上書き対象となる
			indexList := git.SearchIndexList("", pathList)
			overwriteList := git.SearchWorktreeList("", indexList)

			// worktree への上書きがある場合は、バックアップを行う
			if len(overwriteList) != 0 {
				service.Snap(`"fit change unstage" のバックアップ`, pathList...)
				fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
					`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
				exitCode := git.RestoreWorktree(overwriteList)
				if exitCode != 0 {
					return errors.New("restore index failed")
				}
			}
			git.RestoreIndex(pathList)
			git.ShowStatus()
			return nil
		}
	},
	ValidArgs: service.GetStagingFileNameList(),
}

var unstageFlag struct {
	details bool
}

func init() {
	UnstageCmd.Flags().BoolVarP(&unstageFlag.details, "details", "d", false, "ファイルの変更の詳細を表示する")
}
