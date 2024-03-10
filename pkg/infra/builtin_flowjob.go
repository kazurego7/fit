package infra

import "github.com/kazurego7/fit/pkg/domain"

func NewFlowJobCreateCommit() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-createCommit",
		"コミットを作成する",
		func() bool {
			// 変更がある場合
			existsChanges := git.ExistsChanges([]string{":/"})
			return existsChanges
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
		func() bool {
			// 変更がなく、現在のブランチがメインラインでなく、メインライン上にもいない場合
			existsChanges := git.ExistsChanges([]string{":/"})
			currentIsMainline := git.ShowCurrentBranch() == git.GetFitConfig(domain.FitSetting.MainlineType())
			currentIsOnMainline := git.HasContainsCommitOnBranch(git.ShowCurrentBranch(), git.GetFitConfig(domain.FitSetting.MainlineType()))
			return !existsChanges && !currentIsMainline && !currentIsOnMainline
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
		func() bool {
			// 現在のブランチがメインラインでなく、リモートにも存在するが、ローカルとリモートに差分がある場合
			currentBranch := git.ShowCurrentBranch()
			currentIsMainline := currentBranch == git.GetFitConfig(domain.FitSetting.MainlineType())
			remoteIsGone := git.GetUpstreamBranch(currentBranch) == "[gone]"
			currentRemoteCommitId := git.GetShortCommitId(git.GetUpstreamBranch(currentBranch))
			existsDiffLocalToRemote := !git.ExistsUpstreamFor(currentBranch) || git.GetHeadShortCommitId() != currentRemoteCommitId
			return !currentIsMainline && !remoteIsGone && existsDiffLocalToRemote
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
		func() bool {
			// 現在のブランチがメインラインでなく、変更がない場合
			currentBranch := git.ShowCurrentBranch()
			currentIsMainline := currentBranch == git.GetFitConfig(domain.FitSetting.MainlineType())
			existsChanges := git.ExistsChanges([]string{":/"})
			return !currentIsMainline && !existsChanges
		},
		[]domain.FlowTask{},
	)
}

func NewFlowJobSwitchMainline() domain.FlowJob {
	return domain.NewFlowJob(
		"flowJob-switchMainline",
		"メインラインに切り替える",
		func() bool {
			// 現在のブランチがメインラインでなく、リモートにも存在しない場合
			currentBranch := git.ShowCurrentBranch()
			currentIsMainline := currentBranch == git.GetFitConfig(domain.FitSetting.MainlineType())
			remoteIsGone := git.GetUpstreamBranch(currentBranch) == "[gone]"
			return !currentIsMainline && remoteIsGone
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
		func() bool {
			return true
		},
		[]domain.FlowTask{
			NewFlowTaskUpdateMainline(),
		},
	)
}
