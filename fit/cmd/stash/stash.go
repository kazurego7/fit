package stash

import (
	"github.com/spf13/cobra"
)

var StashCmd = &cobra.Command{
	Use:   "stash",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {

	StashCmd.AddCommand(DeleteCmd)
	StashCmd.AddCommand(ListCmd)
	StashCmd.AddCommand(RestoreCmd)
	StashCmd.AddCommand(SnapCmd)
	StashCmd.AddCommand(ShowCmd)
	StashCmd.AddCommand(StoreCmd)
}
