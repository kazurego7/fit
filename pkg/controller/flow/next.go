package flow

import (
	"fmt"
	"strconv"

	"github.com/kazurego7/fit/pkg/infra"
	"github.com/kazurego7/fit/pkg/util"
	"github.com/spf13/cobra"
)

var NextCmd = &cobra.Command{
	Use:   "next",
	Short: "指定したジョブに属するタスクを順に実行する.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ワークフローの提案を取得
		workflow := infra.NewGitlabFlow().FilterExecutable()
		// 引数のバリデーション
		jobNo := 1
		if len(args) != 0 {
			selectedJobNo, err := strconv.Atoi(args[0])
			if err != nil || !workflow.ExistsFlowJobByNo(selectedJobNo) {
				fmt.Println("指定したジョブが存在しません")
				return
			}
			jobNo = selectedJobNo
		}
		flowJob, _ := workflow.FlowJobByNo(jobNo)

		// 実行するジョブの表示
		fmt.Println("【ジョブの開始】")
		fmt.Println("--------------------------------------------------")
		fmt.Println(flowJobFormat(flowJob))
		fmt.Print(flowTaskFormat(flowJob))
		fmt.Println("--------------------------------------------------")
		fmt.Println("上記のジョブを開始しますか？ (Enter or No)")
		isJobExecute, err := util.InputEnterOrNo(false)
		if err != nil || !isJobExecute {
			fmt.Println("ジョブを中止しました")
			return
		}
		// 各タスクの実行
		firstTaskNo, err := flowJob.NextRunTaskNo()
		if err != nil {
			fmt.Println("実行可能なタスクが存在しません")
			return
		}
		taskSequence := flowJob.TaskSequence()[firstTaskNo-1:]
		for index, flowTask := range taskSequence {
			nextTaskNo := index + 1
			if nextTaskNo != 1 {
				fmt.Println("--------------------------------------------------")
				fmt.Println(flowJobFormat(flowJob))
				fmt.Print(flowTaskFormat(flowJob))
				fmt.Println("--------------------------------------------------")
				fmt.Println("上記のタスクを開始しますか？ (Enter or No)")
				isTaskExecute, err := util.InputEnterOrNo(false)
				if err != nil || !isTaskExecute {
					fmt.Println("タスクを中止しました")
					return
				}
			}
			result := flowTask.Run()
			if result.IsFailed() {
				fmt.Println("タスクの実行途中でエラーが発生したため、ジョブを中止しました")
				return
			}
		}
		fmt.Println("**************************************************")
		fmt.Println("*           ジョブの実行が完了しました           *")
		fmt.Print("**************************************************\n\n\n")

		// ワークフローの提案を再取得
		workflow = infra.NewGitlabFlow().FilterExecutable()
		showExecutableFlowjob(workflow, false, 1)
	},
}
