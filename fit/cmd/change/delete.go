package change

import (
	"fmt"
	"os"

	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "ワークツリー・インデックスの変更を削除する.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case deleteFlag.worktree && deleteFlag.index || !deleteFlag.worktree && !deleteFlag.index:
			if !existsUntrackedFiles(args...) && !existsWorktreeDiff(args...) && !existsIndexDiff(args...) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			confirmBackup()
			deleteAll(args...)
		case deleteFlag.worktree:
			if !existsUntrackedFiles(args...) && !existsWorktreeDiff(args...) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			confirmBackup()
			deleteWorktree(args...)
		case deleteFlag.index:
			if !existsIndexDiff(args...) {
				fmt.Fprintln(os.Stderr, "削除するファイルがありません")
				return
			}
			confirmBackup()
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

func existsUntrackedFiles(filenameList ...string) bool {
	if len(filenameList) == 0 {
		return false
	}
	gitSubCmd := append([]string{"ls-files", "--others", "--"}, filenameList...)
	out, _, _ := util.GitQuery(gitSubCmd...)
	list := util.SplitLn(string(out))

	return len(list) != 0
}

func existsWorktreeDiff(args ...string) bool {
	list := searchWorktreeList("", args[0])
	return len(list) != 0
}

func existsIndexDiff(args ...string) bool {
	list := searchIndexList("", args[0])
	return len(list) != 0
}

func deleteWorktree(args ...string) {
	unergedList := searchWorktreeList("U", args[0])
	for i := range unergedList {
		unergedList[i] = ":!" + unergedList[i]
	}
	restoreList := searchWorktreeList("", append(unergedList, args[0])...)
	if len(restoreList) != 0 {
		exitCode := restoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	addedList := searchWorktreeList("A", args[0])
	if len(addedList) != 0 {
		exitCode := removeIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	clean(args[0])
}

func deleteIndex(args ...string) {
	indexList := searchIndexList("", args[0])
	worktreeList := searchWorktreeList("", args[0])
	indexOnlyList := util.Difference(indexList, worktreeList)
	restoreList := searchIndexList("a", indexOnlyList...)
	cleanList := searchIndexList("A", indexOnlyList...)
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
	indexList := searchIndexList("", args[0])
	if len(indexList) != 0 {
		exitCode := restoreIndex(indexList...)
		if exitCode != 0 {
			return
		}
	}
	addedList := searchWorktreeList("A", args[0])
	if len(addedList) != 0 {
		exitCode := removeIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	restoreList := searchWorktreeList("a", args[0])
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
