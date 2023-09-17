package branch

import (
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "ブランチに関する操作.",
}

func init() {
	BranchCmd.AddCommand(ListCmd)
	BranchCmd.AddCommand(SwitchCmd)
	BranchCmd.AddCommand(CreateCmd)
	BranchCmd.AddCommand(DeleteCmd)
	BranchCmd.AddCommand(RenameCmd)
	BranchCmd.AddCommand(FollowCmd)
	BranchCmd.AddCommand(UploadCmd)
}
