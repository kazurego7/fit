package branch

import (
	"strings"

	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Operations on the branch that are temporary and moving marker of a git revision",
}

func init() {
	BranchCmd.AddCommand(CreateCmd)
	BranchCmd.AddCommand(DeleteCmd)
	BranchCmd.AddCommand(ListCmd)
	BranchCmd.AddCommand(DownloadCmd)
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
