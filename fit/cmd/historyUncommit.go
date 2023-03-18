package cmd

import (
	"github.com/kazurego7/fit/fit/gitexec"
	"github.com/spf13/cobra"
)

var historyUncommitCmd = &cobra.Command{
	Use:   "uncommit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"reset", "--soft", "HEAD^"}
		gitexec.Git(gitSubCmd...)
	},
}

func init() {
	historyCmd.AddCommand(historyUncommitCmd)

}
