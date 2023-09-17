package change

import (
	"github.com/spf13/cobra"
)

var ChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "ワークツリー・インデックスの変更に関する操作.",
}

func init() {
	ChangeCmd.AddCommand(ListCmd)
	ChangeCmd.AddCommand(StageCmd)
	ChangeCmd.AddCommand(UnstageCmd)
	ChangeCmd.AddCommand(DeleteCmd)
	ChangeCmd.AddCommand(LogCmd)
	ChangeCmd.AddCommand(RestoreCmd)
}
