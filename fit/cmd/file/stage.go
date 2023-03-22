package file

import (
	"strings"

	"github.com/kazurego7/fit/fit/fitio"
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
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// FIXME: 引数が0この時、ファイル名が補完されてしまう
		if len(toComplete) == 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		gitSubCmd := []string{"ls-files", `-m`, "-o"}
		out, err := fitio.ExecuteGitOutput(gitSubCmd...)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		expect := strings.Split(string(out), "\n")
		return expect, cobra.ShellCompDirectiveNoFileComp
	},
}
