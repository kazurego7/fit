package tag

import (
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "タグに関する操作.",
}

func init() {
	TagCmd.AddCommand(ListCmd)
	TagCmd.AddCommand(CreateCmd)
	TagCmd.AddCommand(DeleteCmd)
	TagCmd.AddCommand(UploadCmd)
}
