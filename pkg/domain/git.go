package domain

type Git interface {
	SearchUntrackedFiles(pathspecs []string) []string
	SearchIndexList(diffFilter string, pathspecs []string) []string
	SearchWorktreeList(diffFilter string, pathspecs []string) []string
	ExistsUntrackedFiles(pathspecs []string) bool
	ExistsWorktreeDiff(pathspecs []string) bool
	ExistsIndexDiff(pathspecs []string) bool
	ExistsChanges(pathspecs []string) bool
	GetShortCommitId(gitrevision string) string
	GetHeadShortCommitId() string
	StashPushAll(stashMessage string, files []string) int
	StashPushOnlyWorktree(stashMessage string) int
	StashApply() int
	CreateBranch(branch string) int
	ShowCurrentBranch() string
	GetUpstreamBranch(localBranch string) string
	ExistsUpstreamFor(branchName string) bool
	GetBranchName(refspec string) string
	PushFor(branch string) int
	PullFor(branch string) int
	SetUpstream(branch string) int
	SwitchBranch(branch string) int
	HasContainsCommitOnBranch(branch string, commit string) bool
	RemoveIndex(filenameList []string) int
	RestoreWorktree(filenameList []string) int
	RestoreIndex(filenameList []string) int
	Clean(filenameList []string) int
	IsConflictResolved(pathspecs []string) bool
	ExistsHEADCommit() bool
	InitRepository()
	FirstCommit()
	ApplyKeepIndex(stashcommit string) int
	Apply(stashcommit string) int
	FetchPrune() int
	GetBranchNameListInUpdateOrder() ([]string, error)
	CommitWithAllowEmpty(message string) int
	CommitWithMessage(message string) int
	CommitWithOpenEditor() int
	RebaseToMainline(branch string) int
	DiffIndex(pathspecList []string) int
	DiffWorktree(pathspecList []string) int
	GetCommitMessage(gitrevision string) string
	StageAll() int
	ResetHeadWithoutWorktree() int
	ResetHeadWithoutWorktreeAndIndex() int
	ShowStatus() int
	ShowChangeDetails() int
	SetConfigDefaultMainline() int
	GetConfig(name string) string
	SetConfig(config Setting) int
}
