package gitImpl

import (
	"strings"

	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/util"
)

type Git struct {
}

func (g Git) addRoot(pathList ...string) []string {

	for i, path := range pathList {
		pathList[i] = ":/" + path
	}
	return pathList
}

func (g Git) SearchUntrackedFiles(pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"ls-files", "--others", "--exclude-standard", "--full-name", "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return g.addRoot(list...)
}

func (g Git) SearchIndexList(diffFilter string, pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--staged", "--no-renames", "--diff-filter=" + diffFilter, "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return g.addRoot(list...)
}

func (g Git) SearchWorktreeList(diffFilter string, pathspecs []string) []string {
	if len(pathspecs) == 0 {
		return []string{}
	}
	gitSubCmd := append([]string{"diff", "--name-only", "--no-renames", "--diff-filter=" + diffFilter, "--"}, pathspecs...)
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	list := util.SplitLn(string(out))
	return g.addRoot(list...)
}

func (g Git) ExistsUntrackedFiles(pathspecs []string) bool {
	list := g.SearchUntrackedFiles(pathspecs)
	return len(list) != 0
}

func (g Git) ExistsWorktreeDiff(pathspecs []string) bool {
	list := g.SearchWorktreeList("", pathspecs)
	return len(list) != 0
}

func (g Git) ExistsIndexDiff(pathspecs []string) bool {
	list := g.SearchIndexList("", pathspecs)
	return len(list) != 0
}

func (g Git) GetHeadShortCommitId() string {
	gitSubCmd := []string{"rev-parse", "--short", "HEAD"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g Git) StashPushAll(stashMessage string, files []string) int {
	gitSubCmd := append([]string{"stash", "push", "--include-untracked", "--"}, files...)
	if stashMessage != "" {
		commitId := g.GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g Git) StashPushOnlyWorktree(stashMessage string) int {
	gitSubCmd := []string{"stash", "push", "--include-untracked", "--keep-index"}
	if stashMessage != "" {
		commitId := g.GetHeadShortCommitId()
		gitSubCmd = append(gitSubCmd, "--message", commitId+" "+stashMessage)
	}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g Git) StashApply() int {
	gitSubCmd := []string{"stash", "apply", "--index", "--quiet"}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g Git) ShowCurrentBranch() string {
	gitSubCmd := []string{"branch", "--show-current"}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g Git) ExistsUpstreamFor(branchName string) bool {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", "--symbolic-full-name", branchName + `@{u}`}
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return exitCode == 0
}

func (g Git) GetBranchName(refspec string) string {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", refspec}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return strings.Trim(string(out), "\n")
}

func (g Git) PullFor(branch string) int {
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

func (g Git) SetUpstream(branch string) int {
	gitSubCmd := []string{"branch", branch, "--set-upstream-to=origin/" + branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g Git) SwitchBranch(branch string) int {
	gitSubCmd := []string{"switch", branch}
	exitCode := util.GitCommand(global.RootFlag, gitSubCmd)
	return exitCode
}

func (g Git) RemoveIndex(filenameList []string) int {
	gitSubCmd := append([]string{"rm", "--cache", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) RestoreWorktree(filenameList []string) int {
	gitSubCmd := append([]string{"restore", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) RestoreIndex(filenameList []string) int {
	gitSubCmd := append([]string{"restore", "--staged", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) Clean(filenameList []string) int {
	gitSubCmd := append([]string{"clean", "--force", "--"}, filenameList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) IsConflictResolved(pathspecs []string) bool {
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

func (g Git) ExistsHEADCommit() bool {
	gitSubCmd := []string{"rev-parse", "HEAD"}
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return exitCode == 0
}

func (g Git) InitRepository() {
	gitSubCmd := []string{"init"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) FirstCommit() {
	gitSubCmd := []string{"commit", "--allow-empty", "-m", "first commit"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) ApplyKeepIndex(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", "--index", stashcommit}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) Apply(stashcommit string) int {
	gitSubCmd := []string{"stash", "apply", "--quiet", stashcommit}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) FetchPrune() {
	gitSubCmd := []string{"fetch", "origin", "--prune"}
	util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) GetBranchNameListInUpdateOrder() ([]string, error) {
	gitSubCmd := []string{"for-each-ref", "--sort=committerdate", `--format="%(refname:lstrip=-1)"`, "refs/remotes", "refs/heads"}
	out, _, err := util.GitQuery(global.RootFlag, gitSubCmd)
	if err != nil {
		return nil, err
	}
	return util.SplitLn(string(out)), err
}

func (g Git) CommitWithAllowEmpty(message string) int {
	// コミット
	gitSubCmd := []string{"commit", "--allow-empty", "--message", message}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) CommitWithMessage(message string) int {
	gitSubCmd := []string{"commit", "--message", message}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) CommitWithOpenEditor() int {
	gitSubCmd := []string{"commit", "--edit"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) DiffIndex(pathspecList []string) int {
	gitSubCmd := append([]string{"diff", "--staged", "--"}, pathspecList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) DiffWorktree(pathspecList []string) int {
	gitSubCmd := append([]string{"diff", "--"}, pathspecList...)
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) GetCommitMessage(gitrevision string) string {
	gitSubCmd := []string{"log", "--format=%B -n 1", gitrevision}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return string(out)
}

func (g Git) AddStageing(pathspecs []string) int {
	gitSubCmd := []string{"add", ":/"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) ResetHeadWithoutWorktree() int {
	gitSubCmd := []string{"reset", "--mixed", "HEAD^"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) ResetHeadWithoutWorktreeAndIndex() int {
	gitSubCmd := []string{"reset", "--soft", "HEAD^"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) ShowStatus() int {
	gitSubCmd := []string{"--paginate", "status", "--short", "--untracked-files=all"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) ShowChangeDetails() int {
	gitSubCmd := []string{"--paginate", "-c", "status.relativePaths=false", "status", "--verbose", "--verbose", "--untracked-files=all"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) SetConfigDefaultMainline() int {
	gitSubCmd := []string{"config", "fit.mainline", "main"}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}

func (g Git) GetFitConfig(name string) string {
	gitSubCmd := []string{"config", "--get", name}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd)
	return util.TrimEOL(string(out))
}

func (g Git) SetFitConfig(config domain.FitConfig) int {
	gitSubCmd := []string{"config", config.GetName(), config.GetValue()}
	return util.GitCommand(global.RootFlag, gitSubCmd)
}
