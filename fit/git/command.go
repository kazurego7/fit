package git

import (
	"errors"
	"strings"

	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

func addRoot(pathList ...string) []string {

	for i, path := range pathList {
		pathList[i] = ":/" + path
	}
	return pathList
}

func SearchUntrackedFiles(filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"ls-files", "--others", "--exclude-standard", "--full-name", "--"}, filenameList...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	list := util.SplitLn(string(out))
	return addRoot(list...)
}

func SearchIndexList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--staged", "--no-renames", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	list := util.SplitLn(string(out))
	return addRoot(list...)
}

func SearchWorktreeList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--no-renames", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	list := util.SplitLn(string(out))
	return addRoot(list...)
}

func ExistsUntrackedFiles(args ...string) bool {
	list := SearchUntrackedFiles(args...)
	return len(list) != 0
}

func ExistsWorktreeDiff(args ...string) bool {
	list := SearchWorktreeList("", args...)
	return len(list) != 0
}

func ExistsIndexDiff(args ...string) bool {
	list := SearchIndexList("", args...)
	return len(list) != 0
}

func GetHeadShortCommitId() string {
	gitSubCmd := []string{"rev-parse", "--short", "HEAD"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return strings.Trim(string(out), "\n")
}

func Snap(stashMessage string, files ...string) int {
	exitCode := StashPushAll(stashMessage, files)
	if exitCode != 0 {
		return exitCode
	}
	return StashApply()
}

func StashPushAll(stashMessage string, files []string) int {
	gitSubCmd := append([]string{"stash", "push", "--include-untracked", "--"}, files...)
	if stashMessage != "" {
		commitId := GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}

func StashPushOnlyWorktree(stashMessage string) int {
	gitSubCmd := []string{"stash", "push", "--include-untracked", "--keep-index"}
	if stashMessage != "" {
		commitId := GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}

func StashApply() int {
	gitSubCmd := []string{"stash", "apply", "--index", "--quiet"}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd...)
	return exitCode
}

func ShowCurrentBranch() string {
	gitSubCmd := []string{"branch", "--show-current"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return strings.Trim(string(out), "\n")
}

func CurrentIsNotReadonly() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if ShowCurrentBranch() == "" {
			return errors.New("現在、読み込み専用の状態です\n" +
				"※ \"fit branch switch\" で特定のブランチに切り替えるか、\"fit branch create\" で新しいブランチに切り替えてください")
		}
		return nil
	}
}
