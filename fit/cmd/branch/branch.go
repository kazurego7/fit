package branch

import (
	"strings"

	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "gitリビジョンを一時的に指し示すマーカーであるブランチに関する操作.",
}

func init() {
	BranchCmd.AddCommand(CreateCmd)
	BranchCmd.AddCommand(DeleteCmd)
	BranchCmd.AddCommand(ListCmd)
	BranchCmd.AddCommand(FollowCmd)
	BranchCmd.AddCommand(UploadCmd)
	BranchCmd.AddCommand(RenameCmd)
}

func existsUpstreamFor(branchName string) bool {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", "--symbolic-full-name", branchName + `@{u}`}
	_, exitCode, _ := util.GitQuery(gitSubCmd...)
	return exitCode == 0
}

func getBranchName(refspec string) string {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", refspec}
	out, _, _ := util.GitQuery(gitSubCmd...)
	return strings.Trim(string(out), "\n")
}
