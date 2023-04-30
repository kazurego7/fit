package revision

import (
	"github.com/spf13/cobra"
)

var RevisionCmd = &cobra.Command{
	Use:   "revision",
	Short: "リビジョンの履歴に関する操作.",
}

func init() {
	RevisionCmd.AddCommand(ListCmd)
	RevisionCmd.AddCommand(CreateCmd)
	RevisionCmd.AddCommand(BackCmd)
	RevisionCmd.AddCommand(DownloadCmd)
	RevisionCmd.AddCommand(MergeCmd)
	RevisionCmd.AddCommand(ShowCmd)
}
