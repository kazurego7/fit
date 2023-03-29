package history

import (
	"github.com/spf13/cobra"
)

var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	HistoryCmd.AddCommand(CommitCmd)
	HistoryCmd.AddCommand(FetchCmd)
	HistoryCmd.AddCommand(GraphCmd)
	HistoryCmd.AddCommand(HeadlogCmd)
	HistoryCmd.AddCommand(MergeCmd)
	HistoryCmd.AddCommand(SwitchCmd)
	HistoryCmd.AddCommand(UncommitCmd)
}
