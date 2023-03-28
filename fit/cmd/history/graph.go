package history

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var GraphCmd = &cobra.Command{
	Use:   "graph",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{
			"log",
			"--graph",
			"--abbrev-commit",
			"--decorate=no",
			"--date=format: %Y-%m-%d %H:%I:%S",
			"--format=format: %C(03)%>|(26)%h%C(reset)   %C(bold 1)%d%C(reset) %C(bold 0)%s%C(reset) %>|(140)%C(reset)  %C(04)%ad%C(reset)  %C(green)%<(16,trunc)%an%C(reset)",
			"--all",
		}
		if allFlag {
			gitSubCmd = append(gitSubCmd, "--reflog")
		}
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}

var allFlag bool

func init() {
	GraphCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "show all commits.")
}
