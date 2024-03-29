package domain

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

type Service struct {
	git Git
}

func NewService(git Git) Service {
	return Service{
		git: git,
	}
}

func (s Service) Snap(stashMessage string, files ...string) int {
	exitCode := s.git.StashPushAll(stashMessage, files)
	if exitCode != 0 {
		return exitCode
	}
	return s.git.StashApply()
}

func (s Service) CurrentIsNotReadonly() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if s.git.ShowCurrentBranch() == "" {
			return errors.New("現在、読み込み専用の状態です\n" +
				"※ \"fit branch switch\" で特定のブランチに切り替えるか、\"fit branch create\" で新しいブランチに切り替えてください")
		}
		return nil
	}
}

func (s Service) ExistsFiles(n int) cobra.PositionalArgs {
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

func (s Service) ExistsWorktreeChanges() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		overwriteList := s.git.SearchWorktreeList("", args)
		if len(overwriteList) != 0 {
			return errors.New("復元するファイルに変更があります.\n" +
				"\"fit change delete\" で変更を削除するか、\"fit change stage\" でステージングを行ってください")
		}
		return nil
	}
}

func (s Service) BackupDelete(pathspecs []string) {
	s.Snap(`"fit change delete" のバックアップ`, pathspecs...)
	fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
		`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
}

func (s Service) DeleteWorktree(pathspecs []string) {
	unergedList := s.git.SearchWorktreeList("U", pathspecs)
	for i := range unergedList {
		unergedList[i] = ":!" + unergedList[i]
	}
	restoreList := s.git.SearchWorktreeList("", append(unergedList, pathspecs...))
	if len(restoreList) != 0 {
		exitCode := s.git.RestoreWorktree(restoreList)
		if exitCode != 0 {
			return
		}
	}
	addedList := s.git.SearchWorktreeList("A", pathspecs)
	if len(addedList) != 0 {
		exitCode := s.git.RemoveIndex(addedList)
		if exitCode != 0 {
			return
		}
	}
	s.git.Clean(pathspecs)
}

func (s Service) DeleteIndex(pathspecs []string) {
	indexList := s.git.SearchIndexList("", pathspecs)
	worktreeList := s.git.SearchWorktreeList("", pathspecs)
	indexOnlyList := util.Difference(indexList, worktreeList)
	restoreList := s.git.SearchIndexList("a", indexOnlyList)
	cleanList := s.git.SearchIndexList("A", indexOnlyList)
	if len(indexList) != 0 {
		exitCode := s.git.RestoreIndex(indexList)
		if exitCode != 0 {
			return
		}
	}
	if len(restoreList) != 0 {
		exitCode := s.git.RestoreWorktree(restoreList)
		if exitCode != 0 {
			return
		}
	}
	if len(cleanList) != 0 {
		exitCode := s.git.Clean(cleanList)
		if exitCode != 0 {
			return
		}
	}
}

func (s Service) DeleteAll(pathspecs []string) {
	indexList := s.git.SearchIndexList("", pathspecs)
	if len(indexList) != 0 {
		exitCode := s.git.RestoreIndex(indexList)
		if exitCode != 0 {
			return
		}
	}
	addedList :=
		s.git.SearchWorktreeList("A", pathspecs)
	if len(addedList) != 0 {
		exitCode := s.git.RemoveIndex(addedList)
		if exitCode != 0 {
			return
		}
	}
	restoreList := s.git.SearchWorktreeList("a", pathspecs)
	if len(restoreList) != 0 {
		exitCode := s.git.RestoreWorktree(restoreList)
		if exitCode != 0 {
			return
		}
	}
	exitCode := s.git.Clean(pathspecs)
	if exitCode != 0 {
		return
	}
}

func (s Service) CheckConflictResolved(pathspecs []string) error {
	isConflictResolved := s.git.IsConflictResolved(pathspecs)
	if isConflictResolved {
		unmergedList := s.git.SearchWorktreeList("U", pathspecs)
		errorMessage := "コンフリクトマーカーが残っています. コンフリクトマーカーを取り除いてください\n" + strings.Join(unmergedList, "\n")
		return errors.New(errorMessage)
	}
	return nil
}

func (s Service) IsBranchOfGone(branch string) bool {
	return s.git.GetUpstreamBranch(branch) == "[gone]"
}

func (s Service) PruneBranchOfGone() {
	// リモートに存在しない上流を持つローカルブランチを取得する
	gitSubCmdGetRefStatus := []string{"for-each-ref", "--format", "%(refname:lstrip=-1):%(upstream:track)"}
	refStatusByte, _, err := util.GitQuery(global.RootFlag, gitSubCmdGetRefStatus)
	currentBranch := s.git.ShowCurrentBranch()
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

func (s Service) StageChange(pathspecList []string) int {
	// index にも worktree にもあるファイルは上書き対象となる
	indexList := s.git.SearchIndexList("u", pathspecList)
	overwriteList := s.git.SearchWorktreeList("", indexList)

	// index への上書きがある場合は、バックアップを行う
	if len(overwriteList) != 0 {
		s.Snap(`"fit change stage" のバックアップ`, pathspecList...)
		fmt.Println("現在のファイルの変更をスタッシュにバックアップしました.\n" +
			`ファイルを復元したい場合は "fit stash restore" を利用してください.`)
	}
	gitSubCmd := append([]string{"add"}, pathspecList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (s Service) GetUnstagingFileNameList() []string {
	pathList := append(s.git.SearchUntrackedFiles([]string{":/"}), s.git.SearchWorktreeList("u", []string{":/"})...)
	filenameList := []string{}
	for _, path := range pathList {
		filename := filepath.Base(path)
		filenameList = append(filenameList, filename)
	}
	return filenameList
}

func (s Service) GetStagingFileNameList() []string {
	pathList := s.git.SearchIndexList("", []string{":/"})
	filenameList := []string{}
	for _, path := range pathList {
		filename := filepath.Base(path)
		filenameList = append(filenameList, filename)
	}
	return filenameList
}

func (s Service) AddFuzzyParentPath(pathList []string) []string {
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

func (s Service) SwitchBranchAfterWIP(branch string) {

	const WIP_MESSAGE = "[WIP]"

	existsChanges := s.git.ExistsIndexDiff([]string{":/"}) || s.git.ExistsUntrackedFiles([]string{":/"}) || s.git.ExistsWorktreeDiff([]string{":/"})
	if existsChanges {
		s.git.CommitWithAllowEmpty(WIP_MESSAGE + " Index")
		s.git.StageAll()
		s.git.CommitWithAllowEmpty(WIP_MESSAGE + " Worktree")
	}

	s.git.SwitchBranch(branch)

	if strings.HasPrefix(s.git.GetCommitMessage("HEAD"), WIP_MESSAGE) {
		s.git.ResetHeadWithoutWorktree()
	}
	if strings.HasPrefix(s.git.GetCommitMessage("HEAD"), WIP_MESSAGE) {
		s.git.ResetHeadWithoutWorktreeAndIndex()
	}
}
