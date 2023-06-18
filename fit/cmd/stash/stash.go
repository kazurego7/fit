package stash

import (
	"github.com/spf13/cobra"
)

var StashCmd = &cobra.Command{
	Use:   "stash",
	Short: "スタッシュに関する操作.",
}

func init() {
	StashCmd.AddCommand(ListCmd)
	StashCmd.AddCommand(StoreCmd)
	StashCmd.AddCommand(RestoreCmd)
	StashCmd.AddCommand(DeleteCmd)
}
