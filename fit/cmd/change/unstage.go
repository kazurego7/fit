package change

import (
	"strings"

	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
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
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := append([]string{"restore", "--staged"}, args...)
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"diff", "--cached", "--name-only"}
		out, _, err := util.GitQuery(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := strings.Split(strings.Trim(strings.ReplaceAll(string(out), `"`, ""), "\n"), "\n")
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
