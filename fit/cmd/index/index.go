package index

import (
	"github.com/spf13/cobra"
)

var IndexCmd = &cobra.Command{
	Use:   "index",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	IndexCmd.AddCommand(DiffCmd)
	IndexCmd.AddCommand(LoadCmd)
	IndexCmd.AddCommand(RestoreCmd)
	IndexCmd.AddCommand(SaveCmd)
	IndexCmd.AddCommand(ListCmd)
	IndexCmd.AddCommand(UnmanageCmd)
}
