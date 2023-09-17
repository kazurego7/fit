package service

import (
	"errors"
	"fit/pkg/infra/git"
	"fit/pkg/util"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Snap(stashMessage string, files ...string) int {
	exitCode := git.StashPushAll(stashMessage, files)
	if exitCode != 0 {
		return exitCode
	}
	return git.StashApply()
}

func CurrentIsNotReadonly() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if git.ShowCurrentBranch() == "" {
			return errors.New("現在、読み込み専用の状態です\n" +
				"※ \"fit branch switch\" で特定のブランチに切り替えるか、\"fit branch create\" で新しいブランチに切り替えてください")
		}
		return nil
	}
}

func ExistsFiles(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(n)(cmd, args); err != nil {
			return err
		}
		for i := 0; i < n; i++ {
			if f, err := os.Stat(args[i]); os.IsNotExist(err) || f.IsDir() {
				return errors.New("ファイルが存在しない、または対象がファイルではありません")
			}
		}
		return nil
	}
}

func ExistsWorktreeChanges() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		overwriteList := git.SearchWorktreeList("", args...)
		if len(overwriteList) != 0 {
			return errors.New("復元するファイルに変更があります.\n" +
				"\"fit change delete\" で変更を削除するか、\"fit change stage\" でステージングを行ってください")
		}
		return nil
	}
}

func BackupDelete(pathspecs []string) {
	Snap(`"fit change delete" のバックアップ`, pathspecs...)
	fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
		`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
}

func DeleteWorktree(pathspecs []string) {
	unergedList := git.SearchWorktreeList("U", pathspecs[0])
	for i := range unergedList {
		unergedList[i] = ":!" + unergedList[i]
	}
	restoreList := git.SearchWorktreeList("", append(unergedList, pathspecs[0])...)
	if len(restoreList) != 0 {
		exitCode := git.RestoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	addedList := git.SearchWorktreeList("A", pathspecs[0])
	if len(addedList) != 0 {
		exitCode := git.RemoveIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	git.Clean(pathspecs[0])
}

func DeleteIndex(pathspecs []string) {
	indexList := git.SearchIndexList("", pathspecs[0])
	worktreeList := git.SearchWorktreeList("", pathspecs[0])
	indexOnlyList := util.Difference(indexList, worktreeList)
	restoreList := git.SearchIndexList("a", indexOnlyList...)
	cleanList := git.SearchIndexList("A", indexOnlyList...)
	if len(indexList) != 0 {
		exitCode := git.RestoreIndex(indexList...)
		if exitCode != 0 {
			return
		}
	}
	if len(restoreList) != 0 {
		exitCode := git.RestoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	if len(cleanList) != 0 {
		exitCode := git.Clean(cleanList...)
		if exitCode != 0 {
			return
		}
	}
}

func DeleteAll(pathspecs []string) {
	indexList := git.SearchIndexList("", pathspecs[0])
	if len(indexList) != 0 {
		exitCode := git.RestoreIndex(indexList...)
		if exitCode != 0 {
			return
		}
	}
	addedList :=
		git.SearchWorktreeList("A", pathspecs[0])
	if len(addedList) != 0 {
		exitCode := git.RemoveIndex(addedList...)
		if exitCode != 0 {
			return
		}
	}
	restoreList := git.SearchWorktreeList("a", pathspecs[0])
	if len(restoreList) != 0 {
		exitCode := git.RestoreWorktree(restoreList...)
		if exitCode != 0 {
			return
		}
	}
	exitCode := git.Clean(pathspecs[0])
	if exitCode != 0 {
		return
	}
}

func CheckConflictResolved(pathspecs []string) error {
	isConflictResolved := git.IsConflictResolved(pathspecs)
	if isConflictResolved {
		unmergedList := git.SearchWorktreeList("U", pathspecs...)
		errorMessage := "コンフリクトマーカーが残っています. コンフリクトマーカーを取り除いてください\n" + strings.Join(unmergedList, "\n")
		return errors.New(errorMessage)
	}
	return nil
}
