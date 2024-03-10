package infra

import (
	"strings"

	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"
	"github.com/samber/lo"
)

var git *gitImpl

func init() {
	git = &gitImpl{}
}

type gitImpl struct {
}

func NewGit() *gitImpl {
	return git
}

func (g gitImpl) addRoot(pathList ...string) []string {

	for i, path := range pathList {
		pathList[i] = ":/" + path
	}
	return pathList
}

func (g gitImpl) SearchUntrackedFiles(pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"ls-files", "--others", "--exclude-standard", "--full-name", "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return g.addRoot(list...)
}

func (g gitImpl) SearchIndexList(diffFilter string, pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--staged", "--no-renames", "--diff-filter=" + diffFilter, "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return g.addRoot(list...)
}

func (g gitImpl) SearchWorktreeList(diffFilter string, pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--no-renames", "--diff-filter=" + diffFilter, "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return g.addRoot(list...)
}

func (g gitImpl) ExistsUntrackedFiles(pathspecs []string) bool {
	list := g.SearchUntrackedFiles(pathspecs)
	return len(list) != 0
}

func (g gitImpl) ExistsWorktreeDiff(pathspecs []string) bool {
	list := g.SearchWorktreeList("", pathspecs)
	return len(list) != 0
}

func (g gitImpl) ExistsIndexDiff(pathspecs []string) bool {
	list := g.SearchIndexList("", pathspecs)
	return len(list) != 0
}

func (g gitImpl) ExistsChanges(pathspecs []string) bool {
	return g.ExistsUntrackedFiles(pathspecs) || g.ExistsWorktreeDiff(pathspecs) || g.ExistsIndexDiff(pathspecs)
}

func (g gitImpl) GetShortCommitId(gitrevision string) string {
	gitSubCmd := []string{"rev-parse", "--short", gitrevision}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g gitImpl) GetHeadShortCommitId() string {
	gitSubCmd := []string{"rev-parse", "--short", "HEAD"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g gitImpl) StashPushAll(stashMessage string, files []string) int {
	gitSubCmd := append([]string{"stash", "push", "--include-untracked", "--"}, files...)
	if stashMessage != "" {
		commitId := g.GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) StashPushOnlyWorktree(stashMessage string) int {
	gitSubCmd := []string{"stash", "push", "--include-untracked", "--keep-index"}
	if stashMessage != "" {
		commitId := g.GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) StashApply() int {
	gitSubCmd := []string{"stash", "apply", "--index", "--quiet"}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) CreateBranch(branch string) int {
	gitSubCmd := []string{"checkout", "-b", branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) ShowCurrentBranch() string {
	gitSubCmd := []string{"branch", "--show-current"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g gitImpl) GetUpstreamBranch(localBranch string) string {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", "--symbolic-full-name", "origin/" + localBranch}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g gitImpl) ExistsUpstreamFor(branchName string) bool {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", "--symbolic-full-name", "origin/" + branchName}
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return exitCode == 0
}

func (g gitImpl) GetBranchName(refspec string) string {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", refspec}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g gitImpl) PushFor(branch string) int {
	gitSubCmd := []string{"push", "origin", branch, "--prune"}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) PullFor(branch string) int {
	// pull したいブランチにチェックアウトしているか、そうでないかで処理を分岐
	currentBranch := g.GetBranchName("HEAD")
	var gitSubCmd []string
	if currentBranch == branch {
		gitSubCmd = []string{"pull", "origin", currentBranch + ":" + currentBranch, "--ff-only", "--prune"}
	} else {
		gitSubCmd = []string{"fetch", "origin", branch + ":" + branch, "--prune"}
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) SetUpstream(branch string) int {
	gitSubCmd := []string{"branch", branch, "--set-upstream-to=origin/" + branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) SwitchBranch(branch string) int {
	gitSubCmd := []string{"switch", branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g gitImpl) HasContainsCommitOnBranch(branch string, commit string) bool {
	gitSubCmd := []string{"branch", `--format="%(refname:short)`, "--contains", commit}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return lo.SomeBy(list, func(s string) bool { return s == branch })
}

func (g gitImpl) RemoveIndex(filenameList []string) int {
	gitSubCmd := append([]string{"rm", "--cache", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) RestoreWorktree(filenameList []string) int {
	gitSubCmd := append([]string{"restore", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) RestoreIndex(filenameList []string) int {
	gitSubCmd := append([]string{"restore", "--staged", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) Clean(filenameList []string) int {
	gitSubCmd := append([]string{"clean", "--force", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) IsConflictResolved(pathspecs []string) bool {
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

func (g gitImpl) ExistsHEADCommit() bool {
	gitSubCmd := []string{"rev-parse", "HEAD"}
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return exitCode == 0
}

func (g gitImpl) InitRepository() {
	gitSubCmd := []string{"init"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) FirstCommit() {
	gitSubCmd := []string{"commit", "--allow-empty", "-m", "first commit"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) ApplyKeepIndex(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", "--index", stashcommit}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) Apply(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", stashcommit}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) FetchPrune() int {
	gitSubCmd := []string{"fetch", "origin", "--prune"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) GetBranchNameListInUpdateOrder() ([]string, error) {
	gitSubCmd := []string{"for-each-ref", "--sort=committerdate", `--format="%(refname:lstrip=-1)"`, "refs/remotes", "refs/heads"}
	out, _, err := util.GitQuery(global.RootFlag, gitSubCmd)
	if err != nil {
		return nil, err
	}
	return util.SplitLn(string(out)), err
}

func (g gitImpl) CommitWithAllowEmpty(message string) int {
	// コミット
	gitSubCmd := []string{"commit", "--allow-empty", "--message", message}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) CommitWithMessage(message string) int {
	gitSubCmd := []string{"commit", "--message", message}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) CommitWithOpenEditor() int {
	gitSubCmd := []string{"commit", "--edit"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) DiffIndex(pathspecList []string) int {
	gitSubCmd := append([]string{"diff", "--staged", "--"}, pathspecList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) DiffWorktree(pathspecList []string) int {
	gitSubCmd := append([]string{"diff", "--"}, pathspecList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) GetCommitMessage(gitrevision string) string {
	gitSubCmd := []string{"log", "--format=%B -n 1", gitrevision}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return string(out)
}

func (g gitImpl) StageAll() int {
	gitSubCmd := []string{"add", ":/"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) ResetHeadWithoutWorktree() int {
	gitSubCmd := []string{"reset", "--mixed", "HEAD^"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) ResetHeadWithoutWorktreeAndIndex() int {
	gitSubCmd := []string{"reset", "--soft", "HEAD^"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) ShowStatus() int {
	gitSubCmd := []string{"--paginate", "status", "--short", "--untracked-files=all"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) ShowChangeDetails() int {
	gitSubCmd := []string{"--paginate", "-c", "status.relativePaths=false", "status", "--verbose", "--verbose", "--untracked-files=all"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) SetConfigDefaultMainline() int {
	gitSubCmd := []string{"config", "fit.mainline", "main"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) GetConfig(name string) string {
	gitSubCmd := []string{"config", "--get", name}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return util.TrimEOL(string(out))
}

func (g gitImpl) SetConfig(config domain.Setting) int {
	gitSubCmd := []string{"config", config.Name(), config.Value()}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g gitImpl) RebaseToMainline(branch string) int {
	gitSubCmd := []string{"rebase", branch}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}
