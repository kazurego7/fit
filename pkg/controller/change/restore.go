package change

import (
	"github.com/kazurego7/fit/pkg/global"
	"github.com/kazurego7/fit/pkg/service"
	"github.com/kazurego7/fit/pkg/util"

	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <pathspec>",
	Short: "コミットに含まれるファイルをワークツリーに復元する.",
	Args:  cobra.MatchAll(service.ExistsFiles(1), service.ExistsWorktreeChanges(), service.CurrentIsNotReadonly()),
	Run: func(cmd *cobra.Command, args []string) {
		gitSubCmd := []string{"restore", "--source", RestoreFlag.commit, args[0]}
		util.GitCommand(global.RootFlag, gitSubCmd)
	},
}

var RestoreFlag struct {
	commit string
	before bool
	after  bool
}

func init() {
	RestoreCmd.Flags().StringVarP(&RestoreFlag.commit, "commit", "c", "HEAD", "コミットを指定する.")
}
