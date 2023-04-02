package conflict

import (
	"github.com/spf13/cobra"
)

var ConflictCmd = &cobra.Command{
	Use:   "conflict",
	Short: "Operations on resolving merge conflicts",
}

func init() {
	ConflictCmd.AddCommand(AbortCmd)
	ConflictCmd.AddCommand(CompleteCmd)
	ConflictCmd.AddCommand(ListCmd)
	ConflictCmd.AddCommand(ResolveCmd)
	ConflictCmd.AddCommand(RestoreCmd)
}
