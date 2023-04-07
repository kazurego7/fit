package conflict

import (
	"github.com/spf13/cobra"
)

var ConflictCmd = &cobra.Command{
	Use:   "conflict",
	Short: "マージコンフリクトの解消に関する操作.",
}

func init() {
	ConflictCmd.AddCommand(AbortCmd)
	ConflictCmd.AddCommand(CompleteCmd)
	ConflictCmd.AddCommand(ListCmd)
	ConflictCmd.AddCommand(ResolveCmd)
	ConflictCmd.AddCommand(RestoreCmd)
}
