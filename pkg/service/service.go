package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/infra/git"
	"github.com/kazurego7/fit/pkg/util"

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
		overwriteList := git.SearchWorktreeList("", args)
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
	unergedList := git.SearchWorktreeList("U", pathspecs)
	for i := range unergedList {
		unergedList[i] = ":!" + unergedList[i]
	}
	restoreList := git.SearchWorktreeList("", append(unergedList, pathspecs...))
	if len(restoreList) != 0 {
		exitCode := git.RestoreWorktree(restoreList)
		if exitCode != 0 {
			return
		}
	}
	addedList := git.SearchWorktreeList("A", pathspecs)
	if len(addedList) != 0 {
		exitCode := git.RemoveIndex(addedList)
		if exitCode != 0 {
			return
		}
	}
	git.Clean(pathspecs)
}

func DeleteIndex(pathspecs []string) {
	indexList := git.SearchIndexList("", pathspecs)
	worktreeList := git.SearchWorktreeList("", pathspecs)
	indexOnlyList := util.Difference(indexList, worktreeList)
	restoreList := git.SearchIndexList("a", indexOnlyList)
	cleanList := git.SearchIndexList("A", indexOnlyList)
	if len(indexList) != 0 {
		exitCode := git.RestoreIndex(indexList)
		if exitCode != 0 {
			return
		}
	}
	if len(restoreList) != 0 {
		exitCode := git.RestoreWorktree(restoreList)
		if exitCode != 0 {
			return
		}
	}
	if len(cleanList) != 0 {
		exitCode := git.Clean(cleanList)
		if exitCode != 0 {
			return
		}
	}
}

func DeleteAll(pathspecs []string) {
	indexList := git.SearchIndexList("", pathspecs)
	if len(indexList) != 0 {
		exitCode := git.RestoreIndex(indexList)
		if exitCode != 0 {
			return
		}
	}
	addedList :=
		git.SearchWorktreeList("A", pathspecs)
	if len(addedList) != 0 {
		exitCode := git.RemoveIndex(addedList)
		if exitCode != 0 {
			return
		}
	}
	restoreList := git.SearchWorktreeList("a", pathspecs)
	if len(restoreList) != 0 {
		exitCode := git.RestoreWorktree(restoreList)
		if exitCode != 0 {
			return
		}
	}
	exitCode := git.Clean(pathspecs)
	if exitCode != 0 {
		return
	}
}

func CheckConflictResolved(pathspecs []string) error {
	isConflictResolved := git.IsConflictResolved(pathspecs)
	if isConflictResolved {
		unmergedList := git.SearchWorktreeList("U", pathspecs)
		errorMessage := "コンフリクトマーカーが残っています. コンフリクトマーカーを取り除いてください\n" + strings.Join(unmergedList, "\n")
		return errors.New(errorMessage)
	}
	return nil
}

func PruneBranchOfGone() {
	// リモートに存在しない上流を持つローカルブランチを取得する
	gitSubCmdGetRefStatus := []string{"for-each-ref", "--format", "%(refname:lstrip=-1):%(upstream:track)"}
	refStatusByte, _, err := util.GitQuery(global.RootFlag, gitSubCmdGetRefStatus)
	currentBranch := git.ShowCurrentBranch()
	noRemoteBranchList := []string{}
	for _, line := range strings.Split(string(refStatusByte), "\n") {
		if len(line) == 0 {
			continue
		}
		items := strings.Split(line, ":")
		// リモートに存在しておらず、かつ現在のブランチでないブランチを選択する
		if items[1] == "[gone]" && items[0] != currentBranch {
			noRemoteBranchList = append(noRemoteBranchList, items[0])
		}
	}
	// ブランチの取得に失敗した場合、またはブランチの取得数が0件の場合、終了する
	if err != nil || len(noRemoteBranchList) == 0 {
		return
	}
	// リモートに存在しない上流を持つローカルブランチを削除する
	gitSubCmdDeleteLocal := append([]string{"branch", "--delete"}, noRemoteBranchList...)
	util.GitCommand(global.RootFlag, gitSubCmdDeleteLocal)
}

func StageChange(pathspecList []string) {
	// index にも worktree にもあるファイルは上書き対象となる
	indexList := git.SearchIndexList("u", pathspecList)
	overwriteList := git.SearchWorktreeList("", indexList)

	// index への上書きがある場合は、バックアップを行う
	if len(overwriteList) != 0 {
		Snap(`"fit change stage" のバックアップ`, pathspecList...)
		fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
			`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
	}
	gitSubCmd := append([]string{"add"}, pathspecList...)
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func GetUnstagingFileNameList() []string {
	pathList := append(git.SearchUntrackedFiles([]string{":/"}), git.SearchWorktreeList("u", []string{":/"})...)
	filenameList := []string{}
	for _, path := range pathList {
		filename := filepath.Base(path)
		filenameList = append(filenameList, filename)
	}
	return filenameList
}

func GetStagingFileNameList() []string {
	pathList := git.SearchIndexList("", []string{":/"})
	filenameList := []string{}
	for _, path := range pathList {
		filename := filepath.Base(path)
		filenameList = append(filenameList, filename)
	}
	return filenameList
}

func AddFuzzyParentPath(pathList []string) []string {
	fuzzyPathList := []string{}
	for _, path := range pathList {
		newPath := ""
		if strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../") || strings.HasPrefix(path, ":/") {
			newPath = path
		} else {
			newPath = "*" + path
		}
		fuzzyPathList = append(fuzzyPathList, newPath)
	}
	return fuzzyPathList
}

func SwitchBranchAfterWIP(branch string) {

	const WIP_MESSAGE = "[WIP]"

	existsChanges := git.ExistsIndexDiff([]string{":/"}) || git.ExistsUntrackedFiles([]string{":/"}) || git.ExistsWorktreeDiff([]string{":/"})
	if existsChanges {
		git.Commit(WIP_MESSAGE + " Index")
		git.AddStageing([]string{":/"})
		git.Commit(WIP_MESSAGE + " Worktree")
	}

	git.SwitchBranch(branch)

	if strings.HasPrefix(git.GetCommitMessage("HEAD"), WIP_MESSAGE) {
		git.ResetHeadWithoutWorktree()
	}
	if strings.HasPrefix(git.GetCommitMessage("HEAD"), WIP_MESSAGE) {
		git.ResetHeadWithoutWorktreeAndIndex()
	}
}
