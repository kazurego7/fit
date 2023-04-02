package branch

import (
	"strings"

	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
