package tag

import (
	"github.com/kazurego7/fit/fit/global"
	"github.com/kazurego7/fit/fit/util"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload tags to remote repository (WARNING: UPLOADED TAGS CAN NOT BE DELETED)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"push", "origin", "--tags", args[0], "--prune"}
		util.PrintGitCommand(global.Flags.Dryrun, gitSubCmd...)
		util.GitCommand(global.Flags.Dryrun, gitSubCmd...)
	},
}
