package tag

import (
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "固定さた目印としてのタグに関する操作.",
}

func init() {
	TagCmd.AddCommand(CreateCmd)
	TagCmd.AddCommand(DeleteCmd)
	TagCmd.AddCommand(ListCmd)
	TagCmd.AddCommand(UploadCmd)
}
