package revision

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "コミットの改定履歴を表示する.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{
			"log",
			"--all",
			"--graph",
			"--abbrev-commit",
			"--decorate=no",
			"--date=format: %Y-%m-%d %H:%I:%S",
			"--format=format: %C(03)%>|(26)%h%C(reset)   %C(bold 1)%d%C(reset) %C(bold 0)%s%C(reset) %>|(140)%C(reset)  %C(04)%ad%C(reset)  %C(green)%<(16,trunc)%an%C(reset)",
		}
		if allFlag {
			gitSubCmd = append(gitSubCmd, "--reflog")
		}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}

var allFlag bool

func init() {
	LogCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "show all commits.")
}
