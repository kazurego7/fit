package remote

import (
	"github.com/spf13/cobra"
)

var RemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	RemoteCmd.AddCommand(FetchCmd)
	RemoteCmd.AddCommand(ListCmd)
	RemoteCmd.AddCommand(SetCmd)
	RemoteCmd.AddCommand(UnsetCmd)
}
