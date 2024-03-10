package flow

import (
	"fmt"

	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra"
	"github.com/spf13/cobra"
)

var FlowCmd = &cobra.Command{
	Use:   "flow",
	Short: "開発のワークフローに関する操作.",
}

var git domain.Git = infra.NewGit()
var service = domain.NewService(git)

func init() {
	FlowCmd.AddCommand(ListCmd)
	FlowCmd.AddCommand(NextCmd)
}

func flowJobFormat(flowJob domain.FlowJob) string {
	return fmt.Sprintf("%2d. %s", flowJob.No(), flowJob.Description())
}

func flowTaskFormat(flowJob domain.FlowJob) string {
	nextRunTaskNo, _ := flowJob.NextRunTaskNo()
	taskList := flowJob.TaskSequence()
	taskSize := len(taskList)
	if taskSize == 0 {
		return "     ** タスクなし\n"
	}
	taskLines := ""
	for index, flowTask := range taskList {
		taskNo := index + 1
		// ツリー表示のための記号を設定
		treeSymbol := "└─"
		if taskNo != taskSize {
			treeSymbol = "├─"
		}
		// マークされたタスクの場合、=> を付ける
		// 完了、またはスキップの場合、○ を付ける
		markPart := "   "
		if taskNo == nextRunTaskNo {
			markPart = "==>"
		} else if taskNo < nextRunTaskNo {
			markPart = " ○ "
		}
		taskLines += fmt.Sprintf(" %s %s %s\n", markPart, treeSymbol, flowTask.Description())
	}
	return taskLines
}
