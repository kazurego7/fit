package history

import (
	"github.com/spf13/cobra"
)

var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Operations on the history of git revisions",
}

func init() {
	HistoryCmd.AddCommand(CommitCmd)
	HistoryCmd.AddCommand(FetchCmd)
	HistoryCmd.AddCommand(GraphCmd)
	HistoryCmd.AddCommand(MergeCmd)
	HistoryCmd.AddCommand(SwitchCmd)
	HistoryCmd.AddCommand(UncommitCmd)
}
