package history

import (
	"github.com/kazurego7/fit/fit/gitexec"
	"github.com/spf13/cobra"
)

var MergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"merge", args[0]}
		gitexec.Git(gitSubCmd...)
	},
}