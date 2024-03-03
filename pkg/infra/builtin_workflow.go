package infra

import "github.com/kazurego7/fit/pkg/domain"

var (
	gitlabFlowJobList = []domain.FlowJob{
		NewFlowJobCreateCommit(),
		NewFlowJobRebaseToMainline(),
		NewFlowJobUploadRemote(),
		NewFlowJobTidyupCommit(),
		NewFlowJobSwitchMainline(),
		NewFlowJobUpdateMainline(),
	}
)

func NewBuiltinGitlabFlow() domain.Workflow {
	return domain.NewWorkflow(gitlabFlowJobList)
}
