package change

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var gitSubCmd []string
		if listFlag.all {
			gitSubCmd = []string{"status", "--short", "--untracked-files=all"}
		} else {
			gitSubCmd = []string{"status", "--short"}
		}

		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}

var listFlag struct {
	all bool
}

func init() {
	ListCmd.Flags().BoolVarP(&listFlag.all, "all", "a", false, "Also shows individual files in untracked directories")
}
