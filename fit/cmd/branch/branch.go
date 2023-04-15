package branch

import (
	"strings"

	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "ブランチに関する操作.",
}

func init() {
	BranchCmd.AddCommand(ListCmd)
	BranchCmd.AddCommand(CreateCmd)
	BranchCmd.AddCommand(DeleteCmd)
	BranchCmd.AddCommand(RenameCmd)
	BranchCmd.AddCommand(FollowCmd)
	BranchCmd.AddCommand(UploadCmd)
}

func existsUpstreamFor(branchName string) bool {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", "--symbolic-full-name", branchName + `@{u}`}
	_, exitCode, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return exitCode == 0
}

func getBranchName(refspec string) string {
	gitSubCmd := []string{"rev-parse", "--abbrev-ref", refspec}
	out, _, _ := util.GitQuery(global.RootFlag, gitSubCmd...)
	return strings.Trim(string(out), "\n")
}
