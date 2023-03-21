package remote

import (
	"github.com/kazurego7/fit/fit/fitio"
	"github.com/spf13/cobra"
)

var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"push", "origin", refspecFlag, "--prune", "--set-upstream"}
		fitio.PrintGitCommand(gitSubCmd...)
		fitio.ExecuteGit(gitSubCmd...)
	},
}

var refspecFlag string

func init() {
	PushCmd.PersistentFlags().StringVar(&refspecFlag, "refspec", "HEAD", `refspec`)
}
