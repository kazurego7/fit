package flow

import (
	"fmt"

	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "現在のフローで実行可能なジョブを一覧表示.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		flowCondition := service.NewFlowJobCondition()
		workflow := infra.NewBuiltinGitlabFlow().FilterExecutable(flowCondition)
		showExecutableFlowjob(workflow, listFlag.all, listFlag.details)
	},
}

var listFlag struct {
	details int
	all     bool
}

func init() {
	ListCmd.Flags().IntVarP(&listFlag.details, "details", "d", 1, "指定したジョブに属するタスクを表示する.")
	ListCmd.Flags().BoolVarP(&listFlag.all, "all", "a", false, "全てのジョブを表示する")
}

func showExecutableFlowjob(workflow domain.Workflow, isAll bool, details int) {
	flowJobList := workflow.ToFlowJobList()
	if len(flowJobList) == 0 {
		fmt.Println("実行可能なジョブが存在しません")
		return
	}

	if !isAll && !workflow.ExistsFlowJobByNo(details) {
		fmt.Println("指定したジョブが存在しません")
		return
	}

	fmt.Println("【現在のフローで実行可能なジョブ・タスク】")
	fmt.Println("--------------------------------------------------")
	for _, flowJob := range flowJobList {
		fmt.Println(flowJobFormat(flowJob))
		if isAll || details == flowJob.No() {
			fmt.Print(flowTaskFormat(flowJob))
		}
	}
	fmt.Println("--------------------------------------------------")
}
