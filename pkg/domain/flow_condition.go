package domain

type FlowJobCondition interface {
	CanCreateCommit() bool
	CanRebaseToMainline() bool
	CanUploadRemote() bool
	CanTidyupCommit() bool
	CanSwitchMainline() bool
	CanUpdateMainline() bool
}

type flowJobCondition struct {
	canCreateCommit     bool
	canRebaseToMainline bool
	canUploadRemote     bool
	canTidyupCommit     bool
	canSwitchMainline   bool
	canUpdateMainline   bool
}

func (f *flowJobCondition) CanCreateCommit() bool {
	return f.canCreateCommit
}

func (f *flowJobCondition) CanRebaseToMainline() bool {
	return f.canRebaseToMainline
}

func (f *flowJobCondition) CanUploadRemote() bool {
	return f.canUploadRemote
}

func (f *flowJobCondition) CanTidyupCommit() bool {
	return f.canTidyupCommit
}

func (f *flowJobCondition) CanSwitchMainline() bool {
	return f.canSwitchMainline
}

func (f *flowJobCondition) CanUpdateMainline() bool {
	return f.canUpdateMainline
}

func (s Service) NewFlowJobCondition() FlowJobCondition {
	currentBranch := s.git.ShowCurrentBranch()
	currentCommitId := s.git.GetHeadShortCommitId()
	mainlineName := s.git.GetFitConfig(FitConfigConstant.MainlineType())
	currentIsOnMainline := s.git.HasContainsCommitOnBranch(currentBranch, mainlineName)
	remoteIsGone := s.IsBranchOfGone(currentBranch)
	currentUpstreamBranch := s.git.GetUpstreamBranch(currentBranch)
	currentRemoteCommitId := s.git.GetShortCommitId(currentUpstreamBranch)
	currentIsMainline := currentBranch == mainlineName
	existsChanges := s.git.ExistsChanges([]string{":/"})
	existsCurrentUpstream := s.git.ExistsUpstreamFor(currentBranch)
	existsDiffLocalToRemote := !existsCurrentUpstream || currentCommitId != currentRemoteCommitId
	flowJobCondition := &flowJobCondition{
		canCreateCommit:     existsChanges,                                                  // 変更がある場合
		canRebaseToMainline: !existsChanges && !currentIsMainline && !currentIsOnMainline,   // 変更がなく、現在のブランチがメインラインでなく、メインライン上にもいない場合
		canUploadRemote:     !currentIsMainline && !remoteIsGone && existsDiffLocalToRemote, // 現在のブランチがメインラインでなく、リモートにも存在するが、ローカルとリモートに差分がある場合
		canTidyupCommit:     !currentIsMainline && !existsChanges,                           // 現在のブランチがメインラインでなく、変更がない場合
		canSwitchMainline:   !currentIsMainline && remoteIsGone,                             // 現在のブランチがメインラインでなく、リモートにも存在しない場合
		canUpdateMainline:   true,                                                           // どの場合でも
	}
	return flowJobCondition
}

var (
	flowTaskResultStatus = FlowTaskResultStatus{
		completed: "completed",
		failed:    "failed",
		skip:      "skip",
	}
)

type FlowTaskResultStatus struct {
	completed string
	failed    string
	skip      string
}

func NewFlowTaskResultCompleted() FlowTaskResult {
	return FlowTaskResult{status: flowTaskResultStatus.completed}
}

func NewFlowTaskResultFailed() FlowTaskResult {
	return FlowTaskResult{status: flowTaskResultStatus.failed}
}

func NewFlowTaskResultSkipped() FlowTaskResult {
	return FlowTaskResult{status: flowTaskResultStatus.skip}
}

type FlowTaskResult struct {
	status string
}

func (f *FlowTaskResult) IsCompleted() bool {
	return f.status == flowTaskResultStatus.completed
}

func (f *FlowTaskResult) IsFailed() bool {
	return f.status == flowTaskResultStatus.failed
}

func (f *FlowTaskResult) IsSkip() bool {
	return f.status == flowTaskResultStatus.skip
}
