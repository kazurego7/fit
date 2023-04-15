package git

import (
	"strings"

	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
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
