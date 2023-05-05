package commit

import (
	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "コミットに関する操作.",
}

func init() {
	CommitCmd.AddCommand(ListCmd)
	CommitCmd.AddCommand(CreateCmd)
	CommitCmd.AddCommand(BackCmd)
	CommitCmd.AddCommand(DownloadCmd)
	CommitCmd.AddCommand(MergeCmd)
	CommitCmd.AddCommand(ShowCmd)
}
