package change

import (
	"fmt"
	"os"

	"github.com/kazurego7/fit/fit/git"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <pathspec>…",
	Short: "ワークツリー・インデックスの変更を削除する.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), git.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case deleteFlag.worktree && deleteFlag.index || !deleteFlag.worktree && !deleteFlag.index:
			if !git.ExistsUntrackedFiles(args...) && !git.ExistsWorktreeDiff(args...) && !git.ExistsIndexDiff(args...) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			backupDelete(args...)
			deleteAll(args...)
		case deleteFlag.worktree:
			if !git.ExistsUntrackedFiles(args...) && !git.ExistsWorktreeDiff(args...) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			backupDelete(args...)
			deleteWorktree(args...)
		case deleteFlag.index:
			if !git.ExistsIndexDiff(args...) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			backupDelete(args...)
			deleteIndex(args...)
		}
	},
}

var deleteFlag struct {
	worktree bool
	index    bool
}

func init() {
	DeleteCmd.Flags().BoolVarP(&deleteFlag.worktree, "worktree", "w", false, "ワークツリーの変更を削除する.")
	DeleteCmd.Flags().BoolVarP(&deleteFlag.index, "index", "i", false, "インデックスの変更を削除する.")
	DeleteCmd.MarkFlagsMutuallyExclusive("worktree", "index")
}

func backupDelete(files ...string) {
	git.Snap(`"fit change delete" のバックアップ`, files...)
	fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
		`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
}

func deleteWorktree(args ...string) {
	unergedList := git.SearchWorktreeList("U", args[0])
	for i := range unergedList {
		unergedList[i] = ":!" + unergedList[i]
	}
	restoreList := git.SearchWorktreeList("", append(unergedList, args[0])...)
	if len(restoreList) != 0 {
		exitCode := restoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	addedList := git.SearchWorktreeList("A", args[0])
	if len(addedList) != 0 {
		exitCode := removeIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	clean(args[0])
}

func deleteIndex(args ...string) {
	indexList := git.SearchIndexList("", args[0])
	worktreeList := git.SearchWorktreeList("", args[0])
	indexOnlyList := util.Difference(indexList, worktreeList)
	restoreList := git.SearchIndexList("a", indexOnlyList...)
	cleanList := git.SearchIndexList("A", indexOnlyList...)
	if len(indexList) != 0 {
		exitCode := restoreIndex(indexList...)
		if exitCode != 0 {
			return
		}
	}
	if len(restoreList) != 0 {
		exitCode := restoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	if len(cleanList) != 0 {
		exitCode := clean(cleanList...)
		if exitCode != 0 {
			return
		}
	}
}

func deleteAll(args ...string) {
	indexList := git.SearchIndexList("", args[0])
	if len(indexList) != 0 {
		exitCode := restoreIndex(indexList...)
		if exitCode != 0 {
			return
		}
	}
	addedList := git.SearchWorktreeList("A", args[0])
	if len(addedList) != 0 {
		exitCode := removeIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	restoreList := git.SearchWorktreeList("a", args[0])
	if len(restoreList) != 0 {
		exitCode := restoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	exitCode := clean(args[0])
	if exitCode != 0 {
		return
	}
}
