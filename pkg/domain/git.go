package domain

type Git interface {
	SearchUntrackedFiles(pathspecs []string) []string
	SearchIndexList(diffFilter string, pathspecs []string) []string
	SearchWorktreeList(diffFilter string, pathspecs []string) []string
	ExistsUntrackedFiles(pathspecs []string) bool
	ExistsWorktreeDiff(pathspecs []string) bool
	ExistsIndexDiff(pathspecs []string) bool
	GetHeadShortCommitId() string
	StashPushAll(stashMessage string, files []string) int
	StashPushOnlyWorktree(stashMessage string) int
	StashApply() int
	ShowCurrentBranch() string
	ExistsUpstreamFor(branchName string) bool
	GetBranchName(refspec string) string
	PullFor(branch string) int
	SetUpstream(branch string) int
	SwitchBranch(branch string) int
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
	FetchPrune()
	GetBranchNameListInUpdateOrder() ([]string, error)
	CommitWithAllowEmpty(message string) int
	CommitWithMessage(message string) int
	CommitWithOpenEditor() int
	DiffIndex(pathspecList []string) int
	DiffWorktree(pathspecList []string) int
	GetCommitMessage(gitrevision string) string
	AddStageing(pathspecs []string) int
	ResetHeadWithoutWorktree() int
	ResetHeadWithoutWorktreeAndIndex() int
	ShowStatus() int
	ShowChangeDetails() int
	SetConfigDefaultMainline() int
	GetFitConfigValue(key string) string
	GetFitConfig(config FitConfig) string
	SetFitConfig(config FitConfig) int
}
