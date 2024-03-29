package infra

import (
	"fmt"

	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/util"
)

func NewFlowTaskCreateFeatureBranch() domain.FlowTask {
	return domain.NewFlowTask(
		"flowTask-createFeatureBranch",
		"フィーチャーブランチを作成する",
		func() bool {
			return git.ShowCurrentBranch() != git.GetConfig(domain.FitConfig().Mainline().Name())
		},
		func() string {
			return "既にフィーチャーブランチが作成されているため、ブランチ作成をスキップします"
		},
		func() domain.FlowTaskResult {
			fmt.Println("ブランチ名を入力してください")
			branchName, err := util.InputTextLn()
			if err != nil {
				return domain.FlowTaskResultStatus.Failed()
			}
			if (git.CreateBranch(branchName)) != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			return domain.FlowTaskResultStatus.Completed()
		},
		nil,
	)
}

func NewFlowTaskStageingAllChanges() domain.FlowTask {
	return domain.NewFlowTask(
		"flowTask-stageingAllChanges",
		"全ての変更をステージングする",
		func() bool {
			return !git.ExistsWorktreeDiff([]string{":/"}) && !git.ExistsUntrackedFiles([]string{":/"})
		},
		func() string {
			return "ステージングする変更がないため、ステージングをスキップします"
		},
		func() domain.FlowTaskResult {
			if git.StageAll() != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			return domain.FlowTaskResultStatus.Completed()
		},
		nil,
	)
}
func NewFlowTaskCreateCommit() domain.FlowTask {
	return domain.NewFlowTask(
		"flowTask-createCommit",
		"コミットを作成する",
		func() bool {
			return false
		},
		nil,
		func() domain.FlowTaskResult {
			if git.CommitWithOpenEditor() != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			return domain.FlowTaskResultStatus.Completed()
		},
		nil,
	)
}

func NewFlowTaskRebaseToMainline() domain.FlowTask {
	return domain.NewFlowTask(
		"flowTask-rebaseToMainline",
		"ブランチをメインラインに付け替える",
		func() bool {
			return false
		},
		nil,
		func() domain.FlowTaskResult {
			mainline := git.GetConfig(domain.FitConfig().Mainline().Name())
			if git.RebaseToMainline(mainline) != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			return domain.FlowTaskResultStatus.Completed()
		},
		nil,
	)
}
func NewFlowTaskUploadRemote() domain.FlowTask {
	return domain.NewFlowTask(
		"flowTask-uploadRemote",
		"リモートリポジトリにブランチとコミットをアップロードする",
		func() bool {
			return false
		},
		nil,
		func() domain.FlowTaskResult {
			if git.PushFor(git.ShowCurrentBranch()) != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			return domain.FlowTaskResultStatus.Completed()
		},
		nil,
	)
}

func NewFlowTaskSwitchMainline() domain.FlowTask {
	return domain.NewFlowTask(
		"flowTask-switchMainline",
		"メインラインに切り替える",
		func() bool {
			return false
		},
		nil,
		func() domain.FlowTaskResult {
			mainline := git.GetConfig(domain.FitConfig().Mainline().Name())
			if git.SwitchBranch(mainline) != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			return domain.FlowTaskResultStatus.Completed()
		},
		nil,
	)
}

func NewFlowTaskUpdateMainline() domain.FlowTask {
	return domain.NewFlowTask(
		"flowTask-updateMainline",
		"メインラインを更新する",
		func() bool {
			return false
		},
		nil,
		func() domain.FlowTaskResult {
			mainline := git.GetConfig(domain.FitConfig().Mainline().Name())
			if git.PullFor(mainline) != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			if git.FetchPrune() != 0 {
				return domain.FlowTaskResultStatus.Failed()
			}
			return domain.FlowTaskResultStatus.Completed()
		},
		nil,
	)
}
