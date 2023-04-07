package history

import (
	"github.com/spf13/cobra"
)

var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "gitリビジョンの履歴に関する操作.",
}

func init() {
	HistoryCmd.AddCommand(CommitCmd)
	HistoryCmd.AddCommand(DownloadCmd)
	HistoryCmd.AddCommand(GraphCmd)
	HistoryCmd.AddCommand(MergeCmd)
	HistoryCmd.AddCommand(SwitchCmd)
	HistoryCmd.AddCommand(UncommitCmd)
}
