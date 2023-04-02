package tag

import (
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Operations on tags, which are fixed and unchanging landmarks",
}

func init() {
	TagCmd.AddCommand(CreateCmd)
	TagCmd.AddCommand(DeleteCmd)
	TagCmd.AddCommand(ListCmd)
	TagCmd.AddCommand(UploadCmd)
}
