package git

import (
	"strings"

	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
)

func SearchIndexList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--relative", "--staged", "--no-renames", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return util.SplitLn(string(out))
}

func SearchWorktreeList(diffFilter string, filenameList ...string) []string {
	if len(filenameList) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--relative", "--no-renames", "--diff-filter=" + diffFilter, "--"}, filenameList...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return util.SplitLn(string(out))
}

func ExistsUntrackedFiles(filenameList ...string) bool {
	if len(filenameList) == 0 {
		return false
	}
	gitSubCmd := append([]string{"ls-files", "--others", "--"}, filenameList...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	list := util.SplitLn(string(out))

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

func GetRepositoryPath() string {
	gitSubCmd := []string{"rev-parse", "--show-toplevel"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return strings.Trim(string(out), "\n")
}

func GetHeadShortCommitId() string {
	gitSubCmd := []string{"rev-parse", "--short", "HEAD"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return strings.Trim(string(out), "\n")
}
