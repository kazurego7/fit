package history

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var GraphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Display the history of git revisions in graph",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{
			"log",
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
	GraphCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "show all commits.")
}
