package infra

import "github.com/kazurego7/fit/pkg/domain"

func NewFlowJobCreateCommit() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-createCommit",
		"コミットを作成する",
		func(condition domain.FlowJobCondition) bool {
			return condition.CanCreateCommit()
		},
		[]domain.FlowTask{
			NewFlowTaskCreateFeatureBranch(),
			NewFlowTaskStageingAllChanges(),
			NewFlowTaskCreateCommit(),
		},
	)
}

func NewFlowJobRebaseToMainline() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-rebaseToMainline",
		"ブランチをメインライン上に付け替える",
		func(condition domain.FlowJobCondition) bool {
			return condition.CanRebaseToMainline()
		},
		[]domain.FlowTask{
			NewFlowTaskRebaseToMainline(),
		},
	)
}

func NewFlowJobUploadRemote() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-uploadRemote",
		"リモートにアップロードする",
		func(condition domain.FlowJobCondition) bool {
			return condition.CanUploadRemote()
		},
		[]domain.FlowTask{
			NewFlowTaskUploadRemote(),
		},
	)
}

func NewFlowJobTidyupCommit() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-tidyupCommit",
		"コミットを整理する",
		func(condition domain.FlowJobCondition) bool {
			return condition.CanTidyupCommit()
		},
		[]domain.FlowTask{},
	)
}

func NewFlowJobSwitchMainline() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-switchMainline",
		"メインラインに切り替える",
		func(condition domain.FlowJobCondition) bool {
			return condition.CanSwitchMainline()
		},
		[]domain.FlowTask{
			NewFlowTaskSwitchMainline(),
			NewFlowTaskUpdateMainline(),
		},
	)
}

func NewFlowJobUpdateMainline() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-updateMainline",
		"メインラインを最新にする",
		func(condition domain.FlowJobCondition) bool {
			return condition.CanUpdateMainline()
		},
		[]domain.FlowTask{
			NewFlowTaskUpdateMainline(),
		},
	)
}
