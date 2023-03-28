package conflict

import (
	"fmt"
	"os"

	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var ResolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkCmd := []string{"-c", "core.whitespace=-trailing-space,-space-before-tab,-indent-with-non-tab,-tab-in-indent,-cr-at-eol", "diff", "--check", args[0]}
		out, _ := fitio.GitQuery(checkCmd...)
		if string(out) != "" {
			fmt.Fprintln(os.Stderr, "conflict markers remain")
			return
		} else {
			gitSubCmd := []string{"add", args[0]}
			allArgs := append(gitSubCmd, args...)
			fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
			fitio.GitCommand(global.Flags.Dryrun, allArgs...)
			return
		}

	},
}
