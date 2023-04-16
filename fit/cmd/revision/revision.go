package revision

import (
	"github.com/spf13/cobra"
)

var RevisionCmd = &cobra.Command{
	Use:   "revision",
	Short: "コミットの改定履歴(リビジョン)に関する操作.",
}

func init() {
	RevisionCmd.AddCommand(ListCmd)
	RevisionCmd.AddCommand(CommitCmd)
	RevisionCmd.AddCommand(UncommitCmd)
	RevisionCmd.AddCommand(DownloadCmd)
	RevisionCmd.AddCommand(MergeCmd)
	RevisionCmd.AddCommand(ShowCmd)
}
