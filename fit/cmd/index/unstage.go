package index

import (
	"strings"

	"github.com/kazurego7/fit/fit/fitio"
	"github.com/spf13/cobra"
)

var UnstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := append([]string{"restore", "--staged"}, args...)
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"diff", "--cached", "--name-only"}
		out, err := fitio.ExecuteGitOutput(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := strings.Split(strings.Trim(strings.ReplaceAll(string(out), `"`, ""), "\n"), "\n")
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
