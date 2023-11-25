package git

import (
	"strings"

	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"
)

func addRoot(pathList ...string) []string {

	for i, path := range pathList {
		pathList[i] = ":/" + path
	}
	return pathList
}

func SearchUntrackedFiles(pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"ls-files", "--others", "--exclude-standard", "--full-name", "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return addRoot(list...)
}

func SearchIndexList(diffFilter string, pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--staged", "--no-renames", "--diff-filter=" + diffFilter, "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return addRoot(list...)
}

func SearchWorktreeList(diffFilter string, pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--no-renames", "--diff-filter=" + diffFilter, "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return addRoot(list...)
}

func ExistsUntrackedFiles(pathspecs []string) bool {
	list := SearchUntrackedFiles(pathspecs)
	return len(list) != 0
}

func ExistsWorktreeDiff(pathspecs []string) bool {
	list := SearchWorktreeList("", pathspecs)
	return len(list) != 0
}

func ExistsIndexDiff(pathspecs []string) bool {
	list := SearchIndexList("", pathspecs)
	return len(list) != 0
}

func GetHeadShortCommitId() string {
	gitSubCmd := []string{"rev-parse", "--short", "HEAD"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func StashPushAll(stashMessage string, files []string) int {
	gitSubCmd := append([]string{"stash", "push", "--include-untracked", "--"}, files...)
	if stashMessage != "" {
		commitId := GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func StashPushOnlyWorktree(stashMessage string) int {
	gitSubCmd := []string{"stash", "push", "--include-untracked", "--keep-index"}
	if stashMessage != "" {
		commitId := GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func StashApply() int {
	gitSubCmd := []string{"stash", "apply", "--index", "--quiet"}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func ShowCurrentBranch() string {
	gitSubCmd := []string{"branch", "--show-current"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func ExistsUpstreamFor(branchName string) bool {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", "--symbolic-full-name", branchName + `@{u}`}
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return exitCode == 0
}

func GetBranchName(refspec string) string {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", refspec}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func PullFor(branch string) int {
	// pull したいブランチにチェックアウトしているか、そうでないかで処理を分岐
	currentBranch := GetBranchName("HEAD")
	var gitSubCmd []string
	if currentBranch == branch {
		gitSubCmd = []string{"pull", "origin", currentBranch + ":" + currentBranch, "--ff-only", "--prune"}
	} else {
		gitSubCmd = []string{"fetch", "origin", branch + ":" + branch, "--prune"}
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func SetUpstream(branch string) int {
	gitSubCmd := []string{"branch", branch, "--set-upstream-to=origin/" + branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func SwitchBranch(branch string) int {
	gitSubCmd := []string{"switch", branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func RemoveIndex(filenameList []string) int {
	gitSubCmd := append([]string{"rm", "--cache", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func RestoreWorktree(filenameList []string) int {
	gitSubCmd := append([]string{"restore", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func RestoreIndex(filenameList []string) int {
	gitSubCmd := append([]string{"restore", "--staged", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func Clean(filenameList []string) int {
	gitSubCmd := append([]string{"clean", "--force", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func IsConflictResolved(pathspecs []string) bool {
	gitSubCmd := append(
		[]string{
			"-c",
			"core.whitespace=-trailing-space,-space-before-tab,-indent-with-non-tab,-tab-in-indent,-cr-at-eol",
			"diff",
			"--check",
		},
		pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return string(out) != ""
}

func ExistsHEADCommit() bool {
	gitSubCmd := []string{"rev-parse", "HEAD"}
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return exitCode == 0
}

func InitGit() {
	gitSubCmd := []string{"init"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func FirstCommit() {
	gitSubCmd := []string{"commit", "--allow-empty", "-m", "first commit"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func ApplyKeepIndex(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", "--index", stashcommit}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func Apply(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", stashcommit}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func FetchPrune() {
	gitSubCmd := []string{"fetch", "origin", "--prune"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func GetBranchNameListInUpdateOrder() ([]string, error) {
	gitSubCmd := []string{"for-each-ref", "--sort=committerdate", `--format="%(refname:lstrip=-1)"`, "refs/remotes", "refs/heads"}
	out, _, err := util.GitQuery(global.RootFlag, gitSubCmd)
	if err != nil {
		return nil, err
	}
	return util.SplitLn(string(out)), err
}

func Commit(message string) int {
	// コミット
	gitSubCmd := []string{"commit", "--allow-empty", "--message", message}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func DiffIndex(pathspecList []string) int {
	gitSubCmd := append([]string{"diff", "--staged", "--"}, pathspecList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func DiffWorktree(pathspecList []string) int {
	gitSubCmd := append([]string{"diff", "--"}, pathspecList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func GetCommitMessage(gitrevision string) string {
	gitSubCmd := []string{"log", "--format=%B -n 1", gitrevision}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return string(out)
}

func AddStageing(pathspecs []string) int {
	gitSubCmd := []string{"add", ":/"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func ResetHeadWithoutWorktree() int {
	gitSubCmd := []string{"reset", "--mixed", "HEAD^"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func ResetHeadWithoutWorktreeAndIndex() int {
	gitSubCmd := []string{"reset", "--soft", "HEAD^"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func ShowStatus() int {
	gitSubCmd := []string{"--paginate", "status", "--short", "--untracked-files=all"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func ShowChangeDetails() int {
	gitSubCmd := []string{"--paginate", "-c", "status.relativePaths=false", "status", "--verbose", "--verbose", "--untracked-files=all"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}
