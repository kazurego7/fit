package history

import (
	"errors"

	"github.com/kazurego7/fit/fit/gitexec"
	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("args error")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"commit", "-m", args[0]}
		gitexec.Git(gitSubCmd...)
	},
}
