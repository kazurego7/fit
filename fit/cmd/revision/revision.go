package revision

import (
	"github.com/spf13/cobra"
)

var RevisionCmd = &cobra.Command{
	Use:   "revision",
	Short: "コミットの改定履歴(リビジョン)に関する操作.",
}

func init() {
	RevisionCmd.AddCommand(CommitCmd)
	RevisionCmd.AddCommand(DownloadCmd)
	RevisionCmd.AddCommand(LogCmd)
	RevisionCmd.AddCommand(MergeCmd)
	RevisionCmd.AddCommand(ShowCmd)
	RevisionCmd.AddCommand(SwitchCmd)
	RevisionCmd.AddCommand(UncommitCmd)
}
