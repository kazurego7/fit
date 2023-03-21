package history

import (
	"strings"

	"github.com/kazurego7/fit/fit/fitio"
	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"switch", args[0]}
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// FIXME: 引数が0この時、ブランチ名でなくファイル名が補完されてしまう
		if len(toComplete) == 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		gitSubCmd := []string{"for-each-ref", `--format="%(refname:short)"`, "refs/remotes", "refs/heads"}
		out, err := fitio.ExecuteGitOutput(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := strings.Split(strings.ReplaceAll(string(out), `"`, ""), "\n")
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
