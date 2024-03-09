package commit

import (
	"github.com/kazurego7/fit/pkg/domain"
	"github.com/kazurego7/fit/pkg/infra"
	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "コミットに関する操作.",
}

var git = infra.Git{}
var service = domain.NewService(git)

func init() {
	CommitCmd.AddCommand(ListCmd)
	CommitCmd.AddCommand(CreateCmd)
	CommitCmd.AddCommand(BackCmd)
	CommitCmd.AddCommand(MergeCmd)
	CommitCmd.AddCommand(SwitchCmd)
	CommitCmd.AddCommand(RebaseCmd)
}
