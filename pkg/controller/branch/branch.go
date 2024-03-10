package branch

import (
	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "ブランチに関する操作.",
}

var git domain.Git = infra.NewGit()
var service = domain.NewService(git)

func init() {
	BranchCmd.AddCommand(ListCmd)
	BranchCmd.AddCommand(SwitchCmd)
	BranchCmd.AddCommand(CreateCmd)
	BranchCmd.AddCommand(DeleteCmd)
	BranchCmd.AddCommand(RenameCmd)
	BranchCmd.AddCommand(FollowCmd)
	BranchCmd.AddCommand(UploadCmd)
}
