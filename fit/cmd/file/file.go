package file

import (
	"github.com/spf13/cobra"
)

var FileCmd = &cobra.Command{
	Use:   "file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	FileCmd.AddCommand(DiffCmd)
	FileCmd.AddCommand(UnstageCmd)
	FileCmd.AddCommand(StageCmd)
	FileCmd.AddCommand(RestoreCmd)
	FileCmd.AddCommand(ListCmd)
}
