package index

import (
	"strings"

	"github.com/kazurego7/fit/fit/fitio"
	"github.com/kazurego7/fit/fit/global"
	"github.com/spf13/cobra"
)

var StageCmd = &cobra.Command{
	Use:   "stage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := append([]string{"add"}, args...)
		fitio.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		fitio.CommandGit(global.Flags.Dryrun, gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		gitSubCmd := []string{"ls-files", `--modified`, "--others"}
		out, err := fitio.QueryGit(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := strings.Split(strings.Trim(strings.ReplaceAll(string(out), `"`, ""), "\n"), "\n")
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
