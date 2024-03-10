package infra

import "github.com/kazurego7/fit/pkg/domain"

func NewGitlabFlow() domain.Workflow {
	return domain.NewWorkflow(
		"gitlab",
		[]domain.FlowJob{
			domain.NewFlowJob(
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
			),
			domain.NewFlowJob(
				"flowJob-rebaseToMainline",
				"ブランチをメインライン上に付け替える",
				func() bool {
					// 変更がなく、現在のブランチがメインラインでなく、メインライン上にもいない場合
					existsChanges := git.ExistsChanges([]string{":/"})
					currentIsMainline := git.ShowCurrentBranch() == git.GetConfig(domain.FitConfig().Mainline().Name())
					currentIsOnMainline := git.HasContainsCommitOnBranch(git.ShowCurrentBranch(), git.GetConfig(domain.FitConfig().Mainline().Name()))
					return !existsChanges && !currentIsMainline && !currentIsOnMainline
				},
				[]domain.FlowTask{
					NewFlowTaskRebaseToMainline(),
				},
			),
			domain.NewFlowJob(
				"flowJob-uploadRemote",
				"リモートにアップロードする",
				func() bool {
					// 現在のブランチがメインラインでなく、リモートにも存在するが、ローカルとリモートに差分がある場合
					currentBranch := git.ShowCurrentBranch()
					currentIsMainline := currentBranch == git.GetConfig(domain.FitConfig().Mainline().Name())
					remoteIsGone := git.GetUpstreamBranch(currentBranch) == "[gone]"
					currentRemoteCommitId := git.GetShortCommitId(git.GetUpstreamBranch(currentBranch))
					existsDiffLocalToRemote := !git.ExistsUpstreamFor(currentBranch) || git.GetHeadShortCommitId() != currentRemoteCommitId
					return !currentIsMainline && !remoteIsGone && existsDiffLocalToRemote
				},
				[]domain.FlowTask{
					NewFlowTaskUploadRemote(),
				},
			),
			domain.NewFlowJob(
				"flowJob-tidyupCommit",
				"コミットを整理する",
				func() bool {
					// 現在のブランチがメインラインでなく、変更がない場合
					currentBranch := git.ShowCurrentBranch()
					currentIsMainline := currentBranch == git.GetConfig(domain.FitConfig().Mainline().Name())
					existsChanges := git.ExistsChanges([]string{":/"})
					return !currentIsMainline && !existsChanges
				},
				[]domain.FlowTask{},
			),
			domain.NewFlowJob(
				"flowJob-switchMainline",
				"メインラインに切り替える",
				func() bool {
					// 現在のブランチがメインラインでなく、リモートにも存在しない場合
					currentBranch := git.ShowCurrentBranch()
					currentIsMainline := currentBranch == git.GetConfig(domain.FitConfig().Mainline().Name())
					remoteIsGone := git.GetUpstreamBranch(currentBranch) == "[gone]"
					return !currentIsMainline && remoteIsGone
				},
				[]domain.FlowTask{
					NewFlowTaskSwitchMainline(),
					NewFlowTaskUpdateMainline(),
				},
			),
			domain.NewFlowJob(
				"flowJob-updateMainline",
				"メインラインを最新にする",
				func() bool {
					return true
				},
				[]domain.FlowTask{
					NewFlowTaskUpdateMainline(),
				},
			),
		})
}
